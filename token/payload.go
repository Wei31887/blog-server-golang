package token

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

var (
	ErrExpiredToken = errors.New("token has expired")
	ErrInvalidToken = errors.New("invalid token")
)

type Payload struct {
	Id       uuid.UUID `json:"id"`
	Username string
	Key      string
	jwt.StandardClaims
}
