package api

import (
	"blog/server/model/response"

	"github.com/gin-gonic/gin"
)

type TagApi struct{}

// TagList : get the list of tag
func (*TagApi) TagList(c *gin.Context) {
	result, err := tagService.TagList()
	if err != nil {
		response.CodeResponse(c, response.ERROR)
		return
	}

	res := response.Response{
		Data: result,
	}
	res.Json(c)
}
