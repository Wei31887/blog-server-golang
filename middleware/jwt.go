package middleware

import (
	G "blog/server/global"
	"blog/server/model/response"
	"blog/server/utils"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// JWT : validation layer between request and response
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.GetHeader("token")

		code := response.SUCCESS
		// Check the token from header is empty or not
		if tokenStr == "" {
			code = response.ERROR_AUTH_CHECK_TOKEN_NOT_FOUND
			res := response.Response{
				Code: code,
				Msg: response.GetMsg(code),
			}
			res.Json(c)
			c.Abort()
			return
		} 
		
		// Record the log
		logger := G.GLOBAL_LOG
		logger.Debug("Header token: ", zap.String("jwt token", tokenStr))

		j := utils.NewJWT()
		tokenCliam, err := j.ParseToken(tokenStr)
		if err != nil {
			code = response.ERROR_AUTH_CHECK_TOKEN_FAIL
		} else if time.Now().Unix() > tokenCliam.ExpiresAt {
			code = response.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
		} 
		
		if code != response.SUCCESS {
			res := response.Response{
				Code: code,
				Msg: response.GetMsg(code),
			}
			res.Json(c)
			c.Abort()
			return
		}

		c.Set("token", tokenCliam)
		c.Next()
	}	
}