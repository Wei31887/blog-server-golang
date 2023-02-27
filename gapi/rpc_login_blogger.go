package gapi

import (
	"context"
	"database/sql"
	"time"

	"blog/server/initialize/global"
	"blog/server/model"
	"blog/server/pb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *GRPCServer) LoginBlogger(ctx context.Context, req *pb.LoginBloggerRequest) (*pb.LoginBloggerResponse, error) {

	blogger := &model.Blogger{
		Username: req.GetUsername(),
		Password: req.GetPassword(),
	}

	resultBlogger, err := bloggerService.FindByName(blogger)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "User not found")
		}
		return nil, status.Errorf(codes.Internal, "Fail to find user")
	}

	if blogger.Password != resultBlogger.Password {
		return nil, status.Errorf(codes.NotFound, "Password incorrect")
	}

	// create access JWT token
	accessToken, accessPayload, err := server.tokenMaker.CreateToken(blogger.Username, global.GLOBAL_CONFIG.JWT.AccessTokenDuration)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "Failed to create access token")
	}

	// create refresh JWT token
	refreshToken, refreshPayload, err := server.tokenMaker.CreateToken(blogger.Username, global.GLOBAL_CONFIG.JWT.RefreshTokenDuration)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "Failed to create access token")
	}

	// store refresh token to session
	session := model.Session{
		Id:           refreshPayload.Id,
		Username:     blogger.Username,
		ClientIp:     "",
		UserAgent:    "",
		RefreshToken: refreshToken,
		ExpiresAt:    refreshPayload.ExpiresAt,
		CreatedAt:    time.Now(),
	}
	err = sessionService.Create(&session)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Fail to create session")
	}

	// send token to client
	res := &pb.LoginBloggerResponse{
		SessionId:             session.Id.String(),
		AccessToken:           accessToken,
		RefreshToken:          refreshToken,
		AccessTokenExpiresAt:  timestamppb.New(accessPayload.ExpiresAt),
		RefreshTokenExpiresAt: timestamppb.New(refreshPayload.ExpiresAt),
	}

	return res, nil
}
