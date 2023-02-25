package service

import (
	G "blog/server/global"
	"blog/server/model"

	"github.com/google/uuid"
)

type SessionService struct {}

func (*SessionService) Create(session *model.Session) (error) {
	return G.GLOBAL_DB.Model(session).Create(session).Error
}

func (*SessionService) GetSession(id uuid.UUID) (model.Session, error) {
	var session model.Session
    err := G.GLOBAL_DB.Where("id = ?", id).First(&session).Error
	return session, err
}