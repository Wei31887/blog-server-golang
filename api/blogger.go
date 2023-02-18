package api

import (
	"blog/server/model/response"

	"github.com/gin-gonic/gin"
)

type BloggerApi struct{}

// FindBlogger : request the blogger information
func (*BloggerApi) FindBlogger(c *gin.Context) {
	result, err := bloggerService.FindIdFirst()
	if err != nil {
		response.MsgResponse(c, response.FORBIDDEN, "blogger is not exist")
		return
	}

	result.Password = ""
	res := response.Response{
		Data: result,
	}
	res.Json(c)
}
