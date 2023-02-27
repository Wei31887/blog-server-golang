package admin

import (
	"blog/server/initialize/global"
	"blog/server/model/response"
	"blog/server/token"
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AdminTokenApi struct{}

type refreshTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}

type refreshTokenResponse struct {
	AccessToken      string       `json:"access_token"`
	AccessExpiredAt  time.Time    `json:"access_expired_at"`
}

func (*AdminTokenApi) RefreshToken(c *gin.Context) {
	var req refreshTokenRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.CodeResponse(c, response.BADREQUEST)
		return
	}

	maker := token.NewJWTMaker(global.GLOBAL_CONFIG.JWT.SigningKey)
	refreshPayload, err := maker.VerifyToken(req.RefreshToken)
	if err != nil {
        response.CodeResponse(c, response.ERROR_AUTH_CHECK_TOKEN_FAIL)
		return
    }

	session, err := sessionService.GetSession(refreshPayload.Id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.CodeResponse(c, response.ERROR_AUTH_CHECK_TOKEN_FAIL)
			return
		}
		response.CodeResponse(c, response.ERROR)
	    return
	}

	if time.Now().After(session.ExpiresAt) {
		response.CodeResponse(c, response.ERROR_AUTH_CHECK_TOKEN_FAIL)
	    return	
	}

	accessToken, accessPayload, err := maker.CreateToken(refreshPayload.Username, global.GLOBAL_CONFIG.JWT.AccessTokenDuration)
	if err != nil {
        response.CodeResponse(c, response.ERROR_AUTH_CHECK_TOKEN_FAIL)
        return
    }

	rsp := refreshTokenResponse{
		AccessToken:      accessToken,
        AccessExpiredAt:  accessPayload.ExpiresAt,
	}

	response.ObjResponse(c, rsp)
}
