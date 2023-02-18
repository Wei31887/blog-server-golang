package admin

import (
	"blog/server/model"
	"blog/server/model/response"
	"blog/server/utils"
	"time"

	"github.com/gin-gonic/gin"
)

type AdminBlogApi struct{}

func (*AdminBlogApi) BlogSave(c *gin.Context) {
	blog := &model.Blog{}
	if err := c.ShouldBindJSON(&blog); err != nil {
		response.CodeResponse(c, response.BADREQUEST)
		return
	}

	code := response.SUCCESS
	if blog.Id <= 0 {
		blog.AddTime = time.Now()
		if err := blogService.Create(blog); err != nil {
			code = response.ERROR
		}
	} else {
		blog.UpdateTime = time.Now()
		if err := blogService.Update(blog); err != nil {
			code = response.ERROR
		}
	}

	response.CodeResponse(c, code)
}

func (*AdminBlogApi) BlogList(c *gin.Context) {
	var page utils.Page
	if err := c.ShouldBindJSON(&page); err != nil {
		response.CodeResponse(c, response.BADREQUEST)
		return
	}

	blog := model.Blog{}
	page.Total = int(blogService.Count())
	result, err := blogService.FindList(&blog, &page)
	if err != nil {
		response.CodeResponse(c, response.ERROR)
		return
	}

	res := response.Response{
		Data: result,
	}
	res.Json(c)
}

func (*AdminBlogApi) BlogFindOne(c *gin.Context) {
	blog := &model.Blog{}
	if err := c.ShouldBindJSON(&blog); err != nil {
		response.CodeResponse(c, response.BADREQUEST)
		return
	}

	result, err := blogService.FindOne(blog)
	if err != nil {
		response.CodeResponse(c, response.ERROR)
		return
	}

	res := response.Response{
		Data: result,
	}
	res.Json(c)
}

func (*AdminBlogApi) BlogDelete(c *gin.Context) {
	blog := &model.Blog{}
	if err := c.ShouldBindJSON(&blog); err != nil {
		response.CodeResponse(c, response.BADREQUEST)
		return
	}

	if err := blogService.Delete(blog); err != nil {
		response.CodeResponse(c, response.ERROR)
		return
	}

	response.CodeResponse(c, response.SUCCESS)
}
