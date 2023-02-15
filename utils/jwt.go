package utils

import (
	G "blog/server/global"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/net/context"
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
	dr, _ := ParseDuration(G.GLOBAL_CONFIG.JWT.ExpireTime)
	expireTime := time.Now().Add(dr)
	
	// create the key: use the MD5 algorithm
	key := strconv.Itoa(time.Now().Nanosecond())
	var claims = Claims{
		Username: userName,
		Key: Md5(key),
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
		// fmt.Println(err)
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}

// GetJWTBlackList
func(j *JWT) GetJWTBlackList(tokenStr string) string {
	return "jwt_balck_list:" + Md5(tokenStr)
}

// JoinBlackList
func (j *JWT) JoinBlackList(tokenStr string) (err error) {
	joinUnix := time.Now().Unix()
	timer := time.Duration(15) * time.Minute
	err = G.GLOBAL_REDIS.SetNX(context.Background(), j.GetJWTBlackList(tokenStr), joinUnix, timer).Err()

	return
}

// IsInBlackList
func(j *JWT) IsInBlackList(tokenStr string) bool {
	valUnixStr, err := G.GLOBAL_REDIS.Get(context.Background(), j.GetJWTBlackList(tokenStr)).Result()
	if err != nil || valUnixStr == "" {
		return false
	}
	valUnix, err := strconv.ParseInt(valUnixStr, 10, 64)
	if err != nil {
		return false
	}
	if time.Now().Unix() - valUnix > G.GLOBAL_CONFIG.JWT.JwtBlacklistGracePeriod {
		return false
	}
	return true
}