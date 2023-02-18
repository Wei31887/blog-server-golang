package admin

import (
	"blog/server/model"
	"blog/server/model/response"
	"blog/server/utils"

	"github.com/gin-gonic/gin"
)

type AdminBlogTypeApi struct{}

func (*AdminBlogTypeApi) BlogTypeList(c *gin.Context) {
	var page utils.Page
	var err error
	if err = c.ShouldBindJSON(&page); err != nil {
		response.CodeResponse(c, response.BADREQUEST)
		return
	}

	if page.Total, err = blogTypeService.FindTypeCount(); err != nil {
		response.CodeResponse(c, response.ERROR)
		return
	}

	result, err := blogTypeService.FindTypeList(page)
	if err != nil {
		response.CodeResponse(c, response.ERROR)
		return
	}

	res := response.Response{
		Data:  result,
		Count: page.Total,
	}
	res.Json(c)
}

func (*AdminBlogTypeApi) BlogTypeSave(c *gin.Context) {
	blogType := &model.BlogType{}
	err := c.ShouldBindJSON(blogType)
	if err != nil {
		response.CodeResponse(c, response.BADREQUEST)
		return
	}

	code := response.SUCCESS
	if blogType.Id <= 0 {
		if err := blogTypeService.Create(blogType); err != nil {
			code = response.ERROR
		}
	} else {
		if err := blogTypeService.Update(blogType); err != nil {
			code = response.ERROR
		}
	}

	response.CodeResponse(c, code)
}

func (*AdminBlogTypeApi) BlogTypeOne(c *gin.Context) {
	blogType := &model.BlogType{}
	if err := c.ShouldBindJSON(blogType); err != nil {
		response.CodeResponse(c, response.BADREQUEST)
		return
	}

	result, err := blogTypeService.FindTypeIdOne(blogType)
	if err != nil {
		response.CodeResponse(c, response.ERROR)
		return
	}
	res := response.Response{
		Data: result,
	}
	res.Json(c)
}

// BlogTypeAll
func (*AdminBlogTypeApi) BlogTypeAll(c *gin.Context) {
	result, err := blogTypeService.FindTypeAll()
	if err != nil {
		response.CodeResponse(c, response.ERROR)
		return
	}
	res := response.Response{
		Data: result,
	}
	res.Json(c)
}

// BlogTypeDelete : Query to delete the certain blogtype
func (*AdminBlogTypeApi) BlogTypeDelete(c *gin.Context) {
	blogType := &model.BlogType{}
	if err := c.ShouldBindJSON(&blogType); err != nil {
		response.CodeResponse(c, response.BADREQUEST)
		return
	}

	err := blogTypeService.Delete(blogType)
	if err != nil {
		response.CodeResponse(c, response.ERROR)
		return
	}
	response.CodeResponse(c, response.SUCCESS)
}
