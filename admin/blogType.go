package admin

import (
	"blog/server/model/response"
	"blog/server/service"
	"blog/server/utils"

	"github.com/gin-gonic/gin"
)

//
func BlogTypeList(c *gin.Context) {
	var page utils.Page
	var err error
	if err = c.BindJSON(&page); err != nil {
		response.ResponseWithCode(c, response.INVALID_PARAMS)
		return
	}

	var blogType service.BlogType
	if page.Total, err = blogType.FindTypeCount(); err != nil {
		response.ResponseWithCode(c, response.ERROR)
		return	
	}

	result, err := blogType.FindTypeList(page) 
	if err != nil {
		response.ResponseWithCode(c, response.ERROR)
		return	
	}
	response.SuccessWithDataCount(c, result, page.Total)
}

func BlogTypeSave(c *gin.Context) {
	var blogType service.BlogType
	err := c.BindJSON(&blogType)
	if err != nil {
		res := response.Response{
			Code: response.INVALID_PARAMS,
			Msg: response.GetMsg(response.INVALID_PARAMS),
		}
		res.Json(c)
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

	res := response.Response{
		Code: code,
		Msg: response.GetMsg(code),
	}
	res.Json(c)
}

func BlogTypeOne(c *gin.Context) {
	var blogType service.BlogType
	if err := c.BindJSON(&blogType); err != nil {
		res := response.Response{
			Code: response.INVALID_PARAMS,
			Msg: response.GetMsg(response.INVALID_PARAMS),
		}
		res.Json(c)
		return
	}

	result, err := blogType.FindTypeIdOne()
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

// BlogTypeAll 
func BlogTypeAll(c *gin.Context) {
	var blogType service.BlogType

	result, err := blogType.FindTypeAll()
	if err != nil {
		response.ResponseWithCode(c, response.ERROR)
		return
	}
	response.SuccessWithData(c, result)
}

// BlogTypeDelete : Query to delete the certain blogtype
func BlogTypeDelete(c *gin.Context) {
	var blogType service.BlogType
	if err := c.BindJSON(&blogType); err != nil {
		response.ResponseWithCode(c, response.INVALID_PARAMS)
		return
	}

	err := blogType.Delete()
	if err != nil {
		response.ResponseWithCode(c, response.ERROR)
		return
	}
	response.SuccessResponse(c)
}