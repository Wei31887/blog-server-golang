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
		res := response.Response{
			Code: response.INVALID_PARAMS,
			Msg: response.GetMsg(response.INVALID_PARAMS),
		}
		res.Json(c)
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

	res := response.Response{
		Code: code,
		Msg: response.GetMsg(code),
	}
	res.Json(c)
}


func BlogList(c *gin.Context) {
	var page utils.Page
	if err := c.BindJSON(&page); err != nil {
		response.ResponseWithCode(c, response.INVALID_PARAMS)
		return
	}

	blog := new(service.Blog)
	page.Total = int(blog.Count())
	result, err := blog.FindList(&page) 
	if err != nil {
		response.ResponseWithCode(c, response.ERROR)
		return
	}

	response.SuccessWithData(c, result)
}


func BlogFindOne(c *gin.Context) {
	var blog service.Blog
	if err := c.BindJSON(&blog); err != nil {
		res := response.Response{
			Code: response.INVALID_PARAMS,
			Msg: response.GetMsg(response.INVALID_PARAMS),
		}
		res.Json(c)
		return
	}

	result, err := blog.FindOne()
	if err != nil {
		res := response.Response{
			Code: response.ERROR,
			Msg: response.GetMsg(response.ERROR),
		}
		res.Json(c)
		return
	}

	res := response.Response{
		Code: response.SUCCESS,
		Msg: response.GetMsg(response.SUCCESS),
		Data: result,
	}
	res.Json(c)
	
}

func BlogDelete(c *gin.Context) {
	var blog service.Blog
	if err := c.BindJSON(&blog); err != nil {
		response.ResponseWithCode(c, response.INVALID_PARAMS)
		return
	}

	if err := blog.Delete(); err != nil {
		response.ResponseWithCode(c, response.ERROR)
		return
	}
	response.SuccessResponse(c)
}