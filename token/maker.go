package token

import "time"

type Maker interface {
	CreateToken(userName string, duration time.Duration) (string, *Payload, error)
	VerifyToken(token string) (*Payload, error)
	GetBlackList(payload *Payload) string
	JoinBlackList(payload *Payload) error
	IsInBlackList(payload *Payload) bool
}
