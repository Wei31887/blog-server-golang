package utils

import (
	G "blog/server/global"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWT struct{
	SigningKey	[]byte
}

func NewJWT() *JWT {
	return &JWT{
		SigningKey: []byte(G.GLOBAL_CONFIG.JWT.SigningKey),
	}
}

type Claims struct {
	Username 	string
	Key			string
	jwt.StandardClaims
}

// GenerateToken
func (j *JWT) GenerateToken(userName string) (jwtToken string, err error) {
	nowTime := time.Now()
	dr, _ := ParseDuration(G.GLOBAL_CONFIG.JWT.ExpireTime)
	expireTime := nowTime.Add(dr * time.Hour)

	// create the key: use the MD5 algorithm
	key := strconv.Itoa(time.Now().Nanosecond())
	var claims = Claims{
		Username: userName,
		// Key: Md5(key),
		Key: key,
		StandardClaims : jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer: G.GLOBAL_CONFIG.JWT.Issuer,
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtToken, err = tokenClaims.SignedString(j.SigningKey)
	return 
}

// ParseToken
func (j *JWT) ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}

