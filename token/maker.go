package token

import "time"

type Maker interface {
	CreateToken(userName string, duration time.Duration) (string, *Payload, error)
	VerifyToken(token string) (*Payload, error)
	GetBlackList(token string) string
	JoinBlackList(token string) (err error)
	IsInBlackList(token string) bool
}