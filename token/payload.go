package token

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrExpiredToken = errors.New("token has expired")
	ErrInvalidToken = errors.New("invalid token")
)

type Payload struct {
	Id        uuid.UUID `json:"id"`
	Username  string
	ExpiresAt time.Time `json:"expires_at"`
	IssuedAt  time.Time `json:"issued_at"`
}

func (p *Payload) Valid() error {
	if p.ExpiresAt.Before(time.Now()) {
		return ErrExpiredToken
	}
	return nil
}
