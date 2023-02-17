package api

import (
	"blog/server/model/response"
	"blog/server/service"

	"github.com/gin-gonic/gin"
)

// FindBlogger : request the blogger information
func FindBlogger(c *gin.Context) {
	var blogger service.Blogger
	result , err := blogger.FindIdFirst()
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