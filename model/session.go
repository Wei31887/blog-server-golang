package model

import (
	"time"

	"github.com/google/uuid"
)

type Session struct {
	Id            uuid.UUID `gorm:"id" json:"id"`
	Username      string    `gorm:"user_name" json:"user_name"`
	RefreshToken string    `gorm:"refresh_token" json:"refresh_token"`
	UserAgent    string    `gorm:"user_agent" json:"user_agent"`
	ClientIp     string    `gorm:"client_ip" json:"client_ip"`
	ExpiresAt    time.Time `gorm:"expires_at" json:"expires_at"`
	CreatedAt    time.Time `gorm:"created_at" json:"created_at"`
}

func (Session) TableName() string {
    return "session"
}