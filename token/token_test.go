package token

import (
	"blog/server/utils"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func GenerateToken(t *testing.T, maker Maker) (token string ){
	userName := utils.RandomString(3)
	issueTime := time.Now()
	expiredTime := issueTime.Add(time.Minute)

	token, payload, err := maker.CreateToken(userName, time.Minute)

	require.NoError(t, err)
	require.NotEmpty(t, token)
	require.NotEmpty(t, payload)

	require.WithinDuration(t, expiredTime, payload.ExpiresAt, time.Second)
	require.WithinDuration(t, issueTime, payload.IssuedAt, time.Second)
	return token
}
func TestJWTMAker(t *testing.T) {
	maker := NewJWTMaker(utils.RandomString(6))
	GenerateToken(t, maker)
}

func TestValidJWT(t *testing.T) {
	maker := NewJWTMaker(utils.RandomString(6))
	token := GenerateToken(t, maker)
	payload, err := maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)
}

func TestExpiredJWT(t *testing.T) {
	maker := NewJWTMaker(utils.RandomString(32))

	token, payload, err := maker.CreateToken(utils.RandomString(3), -time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, token)
	require.NotEmpty(t, payload)

	payload, err = maker.VerifyToken(token)
	require.Error(t, err)
	require.ErrorIs(t, err, ErrExpiredToken)
	require.Nil(t, payload)
}

func TestInvalidToken(t *testing.T) {
	maker := NewJWTMaker(utils.RandomString(6))
	token := utils.RandomString(32)

	payload, err := maker.VerifyToken(token)

	require.Error(t, err)
	require.Empty(t, payload)
}