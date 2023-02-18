package admin

import (
	"blog/server/model/response"
	"blog/server/service"
	"blog/server/utils"
	"time"

	"github.com/gin-gonic/gin"
)

type AdminBlogApi struct {}

func (*AdminBlogApi) BlogSave(c *gin.Context) {
	var blog service.Blog
	if err := c.ShouldBindJSON(&blog); err != nil {
		response.CodeResponse(c, response.BADREQUEST)
		return
	}

	code := response.SUCCESS
	if blog.Id <= 0 {
		blog.AddTime = time.Now()
		if err := blog.Create(); err != nil {
			code = response.ERROR
		}
	} else {
		blog.UpdateTime = time.Now()
		if err := blog.Update(); err != nil {
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

	blog := new(service.Blog)
	page.Total = int(blog.Count())
	result, err := blog.FindList(&page) 
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
	var blog service.Blog
	if err := c.ShouldBindJSON(&blog); err != nil {
		response.CodeResponse(c, response.BADREQUEST)
		return
	}

	result, err := blog.FindOne()
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
	var blog service.Blog
	if err := c.ShouldBindJSON(&blog); err != nil {
		response.CodeResponse(c, response.BADREQUEST)
		return
	}

	if err := blog.Delete(); err != nil {
		response.CodeResponse(c, response.ERROR)
		return
	}
	
	response.CodeResponse(c, response.SUCCESS)
}