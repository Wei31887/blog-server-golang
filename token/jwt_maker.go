package token

import (
	"errors"
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
	id, err := uuid.NewRandom()
	if err != nil {
		return "", nil, err
	}

	// create the key: use the MD5 algorithm
	payload := &Payload{
		Id:        id,
		Username:  userName,
		ExpiresAt: time.Now().Add(duration),
		IssuedAt:  time.Now(),
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
