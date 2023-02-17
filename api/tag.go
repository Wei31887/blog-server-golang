package api

import (
	"blog/server/model/response"
	"blog/server/service"

	"github.com/gin-gonic/gin"
)

// TagList : get the list of tag
func TagList(c *gin.Context) {
	var tag service.Tag
	result, err := tag.TagList()
	if err != nil {
		response.CodeResponse(c, response.ERROR)
		return
	}
	
	res := response.Response{
		Data: result,
	}
	res.Json(c)
}