package middleware

import (
	G "blog/server/global"
	"blog/server/model/response"
	"blog/server/token"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// JWT : middleware function, validation layer between request and response
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.GetHeader("token")

		// Check the token from header is empty or not
		if tokenStr == "" {
			response.CodeResponse(c, response.ERROR_AUTH_CHECK_TOKEN_NOT_FOUND)
			c.AbortWithStatus(response.ERROR_AUTH_CHECK_TOKEN_NOT_FOUND)
			return
		} 
		
		// Record the log
		logger := G.GLOBAL_LOG
		logger.Debug("Header token: ", zap.String("jwt token:", tokenStr))
		
		code := response.SUCCESS

		// JWT authentication
		maker := token.NewJWTMaker(G.GLOBAL_CONFIG.JWT.SigningKey)
		// Validate the JWT token and the JWT token is in the black list or not
		payload, err := maker.VerifyToken(tokenStr)
		if err != nil {
			code = response.ERROR_AUTH_CHECK_TOKEN_FAIL 
		} else if maker.IsInBlackList(tokenStr) {
			code = response.ERROR_AUTH_CHECK_TOKEN_IN_BLACK_LIST
		}
		
		if code != response.SUCCESS {
			response.CodeResponse(c, code)
			c.AbortWithStatus(code)
			return
		}

		c.Set("token", payload)
		c.Next()
	}	
}