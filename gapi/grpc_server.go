package gapi

import (
	"blog/server/initialize/config"
	"blog/server/pb"
	"blog/server/token"

	"github.com/gin-gonic/gin"
)

type GRPCServer struct {
	Router *gin.Engine
	Config *config.Config
	pb.UnimplementedBlogServerServer
	tokenMaker token.Maker
}

func NewGRPCServer(config *config.Config) *GRPCServer {
	maker := token.NewJWTMaker(config.JWT.SigningKey)
	server := &GRPCServer{
		Config:     config,
		tokenMaker: maker,
	}
	return server
}
