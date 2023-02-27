package admin

import (
	"blog/server/initialize/global"
	"blog/server/middleware"
	"blog/server/model"
	"blog/server/model/response"
	"blog/server/token"
	"database/sql"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AdminBloggerApi struct{}

type loginResponse struct {
	SessionId        uuid.UUID `json:"session_id"`
	AccessToken      string    `json:"access_token"`
	AccessExpiredAt  time.Time `json:"access_token_expired_at"`
	RefreshToken     string    `json:"refresh_token"`
	RefreshExpiredAt time.Time `json:"refresh_token_expired_at"`
}

// Login
func (*AdminBloggerApi) Login(c *gin.Context) {
	blogger := &model.Blogger{}
	err := c.ShouldBindJSON(blogger)
	if err != nil {
		response.CodeResponse(c, response.BADREQUEST)
		return
	}

	resultBlogger, err := bloggerService.FindByName(blogger)
	if err != nil {
		if err == sql.ErrNoRows {
			response.MsgResponse(c, response.FORBIDDEN, "User not found")
			return
		}
		response.CodeResponse(c, response.ERROR)
		return
	}

	if blogger.Password != resultBlogger.Password {
		response.MsgResponse(c, response.FORBIDDEN, "Wrong password!")
		return
	}

	// create access JWT token
	maker := token.NewJWTMaker(global.GLOBAL_CONFIG.JWT.SigningKey)
	accessToken, accessPayload, err := maker.CreateToken(blogger.Username, global.GLOBAL_CONFIG.JWT.AccessTokenDuration)
	if err != nil {
		response.CodeResponse(c, response.ERROR)
		return
	}

	// create refresh JWT token
	refreshToken, refreshPayload, err := maker.CreateToken(blogger.Username, global.GLOBAL_CONFIG.JWT.RefreshTokenDuration)
	if err != nil {
		response.CodeResponse(c, response.ERROR)
		return
	}

	// store refresh token to session
	session := model.Session{
		Id:           refreshPayload.Id,
		Username:     blogger.Username,
		ClientIp:     c.ClientIP(),
		UserAgent:    c.Request.UserAgent(),
		RefreshToken: refreshToken,
		ExpiresAt:    refreshPayload.ExpiresAt,
		CreatedAt:    time.Now(),
	}
	err = sessionService.Create(&session)
	if err!= nil {
        response.CodeResponse(c, response.ERROR)
        return
    }

	// send token to client
	res := loginResponse{
		SessionId:        session.Id,
		AccessToken:      accessToken,
		AccessExpiredAt:  accessPayload.ExpiresAt,
		RefreshToken:     refreshToken,
		RefreshExpiredAt: refreshPayload.ExpiresAt,
	}
	response.ObjResponse(c, res)
}

// Logout
func (*AdminBloggerApi) Logout(c *gin.Context) {
	maker := token.NewJWTMaker(global.GLOBAL_CONFIG.JWT.SigningKey)
	payload := c.MustGet(middleware.AuthorizationKey).(*token.Payload)

	if valid := maker.IsInBlackList(payload); valid {
		response.CodeResponse(c, response.ERROR_AUTH_CHECK_TOKEN_IN_BLACK_LIST)
		return
	}

	if err := maker.JoinBlackList(payload); err != nil {
		response.CodeResponse(c, response.ERROR)
		return
	}

	response.CodeResponse(c, response.SUCCESS)
}

func (*AdminBloggerApi) FindBlogger(c *gin.Context) {
	queryBlogger, err := bloggerService.FindIdFirst()
	if err != nil {
		response.MsgResponse(c, response.FORBIDDEN, "User not found")
		return
	}
	res := response.Response{
		Data: queryBlogger,
	}
	res.Json(c)
}

// Update blogger password
func (*AdminBloggerApi) UpdatePassword(c *gin.Context) {
	blogger := &model.Blogger{}
	err := c.ShouldBindJSON(blogger)
	if err != nil {
		response.CodeResponse(c, response.BADREQUEST)
		return
	}

	code := response.SUCCESS
	if blogger.Id <= 0 {
		if err = bloggerService.Create(blogger); err != nil {
			code = response.ERROR
		}
	} else {
		if _, err = bloggerService.UpdateSecurityInfo(blogger); err != nil {
			code = response.ERROR
		}
	}

	response.CodeResponse(c, code)
}

func (*AdminBloggerApi) UpdateInfo(c *gin.Context) {
	blogger := &model.Blogger{}
	err := c.ShouldBindJSON(blogger)
	if err != nil {
		response.CodeResponse(c, response.BADREQUEST)
		return
	}

	code := response.SUCCESS
	if blogger.Id <= 0 {
		if err := bloggerService.Create(blogger); err != nil {
			code = response.ERROR
		}
	} else {
		if _, err := bloggerService.UpdateInfo(blogger); err != nil {
			code = response.ERROR
		}
	}

	response.CodeResponse(c, code)
}
