package middleware

import (
	"blog/server/initialize/global"
	"blog/server/model/response"
	"blog/server/token"
	"errors"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	AuthorizationKey = "token"
)

// JWT : middleware function, validation layer between request and response
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.GetHeader("access_token")

		// Check the token from header is empty or not
		if len(tokenStr) == 0 {
			err := errors.New("token is empty")
			c.AbortWithStatusJSON(response.ERROR_AUTH_CHECK_TOKEN_NOT_FOUND, response.ErrResponseJson(err))
			return
		} 
		
		// Record the log
		logger := global.GLOBAL_LOG
		logger.Debug("Header token: ", zap.String("jwt token:", tokenStr))
		
		code := response.SUCCESS

		// JWT authentication
		maker := token.NewJWTMaker(global.GLOBAL_CONFIG.JWT.SigningKey)
		// Validate the JWT token and the JWT token is in the black list or not
		payload, err := maker.VerifyToken(tokenStr)
		if err != nil {
			code = response.ERROR_AUTH_CHECK_TOKEN_FAIL 
		} else if maker.IsInBlackList(payload) {
			code = response.ERROR_AUTH_CHECK_TOKEN_IN_BLACK_LIST
		}
		
		if code != response.SUCCESS {
			response.CodeResponse(c, code)
			c.AbortWithStatus(code)
			return
		}

		c.Set(AuthorizationKey, payload)
		c.Next()
	}	
}