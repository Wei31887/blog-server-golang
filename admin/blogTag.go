package admin

import (
	"blog/server/model"
	"blog/server/model/response"
	"blog/server/utils"

	"github.com/gin-gonic/gin"
)

type AdminBlogTagApi struct {}

func (*AdminBlogTagApi) BlogTagAll(c *gin.Context) {
	result, err := tagService.TagAll()
	if err != nil {
		response.CodeResponse(c, response.ERROR)
		return
	}
	res := response.Response{
		Data: result,
	}
	res.Json(c)
}

func (*AdminBlogTagApi) BlogTagList(c *gin.Context) {
	var page utils.Page
	var err error
	if err = c.ShouldBindJSON(&page); err != nil {
		response.CodeResponse(c, response.BADREQUEST)
		return
	}

	if page.Total, err = tagService.Count(); err != nil {
		response.CodeResponse(c, response.ERROR)
		return
	}

	result, err := tagService.ListPage(page)
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

func (*AdminBlogTagApi) BlogTagSave(c *gin.Context) {
	var tag model.Tag
    if err := c.ShouldBindJSON(&tag); err!= nil {
        response.CodeResponse(c, response.BADREQUEST)
        return
    }

	err := tagService.Create(&tag)
	if err != nil {
		response.CodeResponse(c, response.ERROR)
        return
	}

	response.CodeResponse(c, response.SUCCESS)
}

func (*AdminBlogTagApi) BlogTagDelete(c *gin.Context) {
	var tag model.Tag
    if err := c.ShouldBindJSON(&tag); err!= nil {
        response.CodeResponse(c, response.BADREQUEST)
        return
    }

	err := tagService.Delete(&tag)
	if err != nil {
		response.CodeResponse(c, response.ERROR)
        return
	}

	response.CodeResponse(c, response.SUCCESS)	
}