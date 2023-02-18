package token

import (
	G "blog/server/global"
	"blog/server/utils"
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

type JWTMaker struct {
	SigningKey []byte
}

func NewJWTMaker(signingKey string) Maker {
	return &JWTMaker{
		SigningKey: []byte(signingKey),
	}
}

// GenerateToken
func (maker *JWTMaker) CreateToken(userName string, duration time.Duration) (string, *Payload, error) {
	expireTime := time.Now().Add(duration)
	id, err := uuid.NewRandom()
	if err != nil {
		return "", nil, err
	}

	// create the key: use the MD5 algorithm
	key := strconv.Itoa(time.Now().Nanosecond())
	payload := &Payload{
		Id:       id,
		Username: userName,
		Key:      utils.Md5(key),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    G.GLOBAL_CONFIG.JWT.Issuer,
		},
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	token, err := jwtToken.SignedString(maker.SigningKey)
	if err != nil {
		return "", nil, err
	}
	return token, payload, nil
}

// ParseToken
func (maker *JWTMaker) VerifyToken(token string) (*Payload, error) {
	keyFunc := func(jwtToken *jwt.Token) (interface{}, error) {
		_, ok := jwtToken.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidToken
		}
		return maker.SigningKey, nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, keyFunc)
	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, ErrExpiredToken) {
			return nil, ErrExpiredToken
		}
		return nil, ErrInvalidToken
	}

	payload, ok := jwtToken.Claims.(*Payload)
	if !ok {
		return nil, ErrInvalidToken
	}

	return payload, nil
}

// GetJWTBlackList
func (maker *JWTMaker) GetBlackList(token string) string {
	return "jwt_balck_list:" + utils.Md5(token)
}

// JoinBlackList
func (maker *JWTMaker) JoinBlackList(token string) error {
	joinUnix := time.Now().Unix()
	timer := time.Duration(G.GLOBAL_CONFIG.JWT.JwtBlacklistGracePeriod) * time.Minute
	err := G.GLOBAL_REDIS.SetNX(context.Background(), maker.GetBlackList(token), joinUnix, timer).Err()
	return err
}

// IsInBlackList
func (maker *JWTMaker) IsInBlackList(token string) bool {
	valUnixStr, err := G.GLOBAL_REDIS.Get(context.Background(), maker.GetBlackList(token)).Result()
	if err != nil || valUnixStr == "" {
		return false
	}
	valUnix, err := strconv.ParseInt(valUnixStr, 10, 64)
	if err != nil {
		return false
	}
	if time.Now().Unix()-valUnix > G.GLOBAL_CONFIG.JWT.JwtBlacklistGracePeriod {
		return false
	}
	return true
}
