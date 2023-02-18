package admin

import (
	"blog/server/model/response"
	"blog/server/service"
	"blog/server/utils"

	"github.com/gin-gonic/gin"
)

type AdminBlogTypeApi struct {}

//
func (*AdminBlogTypeApi) BlogTypeList(c *gin.Context) {
	var page utils.Page
	var err error
	if err = c.ShouldBindJSON(&page); err != nil {
		response.CodeResponse(c, response.BADREQUEST)
		return
	}

	var blogType service.BlogType
	if page.Total, err = blogType.FindTypeCount(); err != nil {
		response.CodeResponse(c, response.ERROR)
		return	
	}

	result, err := blogType.FindTypeList(page) 
	if err != nil {
		response.CodeResponse(c, response.ERROR)
		return	
	}

	res := response.Response{
		Data: result,
		Count: page.Total,
	}
	res.Json(c)
}

func (*AdminBlogTypeApi) BlogTypeSave(c *gin.Context) {
	var blogType service.BlogType
	err := c.ShouldBindJSON(&blogType)
	if err != nil {
		response.CodeResponse(c, response.BADREQUEST)
		return	
	}

	code := response.SUCCESS
	if blogType.Id <= 0 {
		if err := blogType.Create(); err != nil {
			code = response.ERROR
		}
	} else {
		if err := blogType.Update(); err != nil {
			code = response.ERROR
		}
	}

	response.CodeResponse(c, code)
}

func (*AdminBlogTypeApi) BlogTypeOne(c *gin.Context) {
	var blogType service.BlogType
	if err := c.ShouldBindJSON(&blogType); err != nil {
		response.CodeResponse(c, response.BADREQUEST)
		return
	}

	result, err := blogType.FindTypeIdOne()
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
	var blogType service.BlogType

	result, err := blogType.FindTypeAll()
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
	var blogType service.BlogType
	if err := c.ShouldBindJSON(&blogType); err != nil {
		response.CodeResponse(c, response.BADREQUEST)
		return
	}

	err := blogType.Delete()
	if err != nil {
		response.CodeResponse(c, response.ERROR)
		return
	}
	response.CodeResponse(c, response.SUCCESS)
}