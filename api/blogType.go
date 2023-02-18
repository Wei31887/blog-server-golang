package api

import (
	"blog/server/model/response"

	"github.com/gin-gonic/gin"
)

type BlogTypeApi struct {}

// FindType : request the amount of each type
func (*BlogTypeApi) FindType(c *gin.Context) {
	result, err := blogTypeService.FindAllTypeCount()
	
	if err != nil {
		response.CodeResponse(c, response.BADREQUEST)
		return
	}
	res := response.Response{
		Data: result,
	}
	res.Json(c)
}