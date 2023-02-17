package admin

import (
	"blog/server/model/response"
	"blog/server/service"
	"blog/server/utils"
	"time"

	G "blog/server/global"

	"github.com/gin-gonic/gin"
)

func BlogSave(c *gin.Context) {
	var blog service.Blog
	if err := c.BindJSON(&blog); err != nil {
		response.CodeResponse(c, response.BADREQUEST)
		return
	}

	code := response.SUCCESS
	if blog.Id <= 0 {
		blog.AddTime = time.Now().Format(G.DateFormat)
		if err := blog.Create(); err != nil {
			code = response.ERROR
		}
	} else {
		blog.UpdateTime = time.Now().Format(G.DateFormat)
		if err := blog.Update(); err != nil {
			code = response.ERROR
		}
	}

	response.CodeResponse(c, code)
}


func BlogList(c *gin.Context) {
	var page utils.Page
	if err := c.BindJSON(&page); err != nil {
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


func BlogFindOne(c *gin.Context) {
	var blog service.Blog
	if err := c.BindJSON(&blog); err != nil {
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

func BlogDelete(c *gin.Context) {
	var blog service.Blog
	if err := c.BindJSON(&blog); err != nil {
		response.CodeResponse(c, response.BADREQUEST)
		return
	}

	if err := blog.Delete(); err != nil {
		response.CodeResponse(c, response.ERROR)
		return
	}
	
	response.CodeResponse(c, response.SUCCESS)
}