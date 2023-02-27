package service

import (
	"blog/server/initialize/global"
	"blog/server/model"

	"github.com/google/uuid"
)

type SessionService struct {}

func (*SessionService) Create(session *model.Session) (error) {
	return global.GLOBAL_DB.Model(session).Create(session).Error
}

func (*SessionService) GetSession(id uuid.UUID) (model.Session, error) {
	var session model.Session
    err := global.GLOBAL_DB.Where("id = ?", id).First(&session).Error
	return session, err
}