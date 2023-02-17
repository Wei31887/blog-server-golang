package admin

import (
	"blog/server/model/response"
	"blog/server/service"
	"blog/server/utils"

	"github.com/gin-gonic/gin"
)

// Login
func Login(c *gin.Context) {
	var blogger service.Blogger
	err := c.BindJSON(&blogger)
	if err != nil {
		response.CodeResponse(c, response.BADREQUEST)
		return
	}

	queryBloger, _ := blogger.FindByName()
	if queryBloger == nil {
		response.MsgResponse(c, response.FORBIDDEN, "User not found")
		return
	}

	if blogger.Password != queryBloger.Password {
		response.MsgResponse(c, response.FORBIDDEN, "Wrong password!")
		return
	}

	// create JWT token 
	j := utils.NewJWT()
	jwtToken, err := j.GenerateToken(blogger.Username)
	if err != nil {
		response.CodeResponse(c, response.ERROR)
		return
	}

	res := response.Response{
		Data: jwtToken,
	}
	res.Json(c)
}

// Logout
func Logout(c *gin.Context) {
	j := utils.NewJWT()
	if err := j.JoinBlackList(c.GetHeader("token")); err != nil {
		response.CodeResponse(c, response.ERROR_AUTH_CHECK_TOKEN_IN_BLACK_LIST)
		return
	}
	response.CodeResponse(c, response.SUCCESS)
}

//
func FindBlogger(c *gin.Context) {
	var blogger service.Blogger
	queryBlogger, err := blogger.FindIdFirst()
	if err != nil {
		response.MsgResponse(c, response.FORBIDDEN, "User not found")
		return
	}
	res := response.Response{
		Data: queryBlogger,
	}
	res.Json(c)
}

// Update blogger password
func UpdatePassword(c *gin.Context) {
	var blogger service.Blogger
	err := c.BindJSON(&blogger)
	if err != nil {
		response.CodeResponse(c, response.BADREQUEST)
		return
	}

	code := response.SUCCESS
	if blogger.Id <= 0 {
		if err = blogger.Create(); err != nil {
			code = response.ERROR
		}
	} else {
		if err = blogger.UpdatePassword(); err != nil {
			code = response.ERROR
		}
	}

	response.CodeResponse(c, code)
}

func UpdateInfo(c *gin.Context) {
	var blogger service.Blogger
	err := c.BindJSON(&blogger)
	if err != nil {
		response.CodeResponse(c, response.BADREQUEST)
		return
	}

	code := response.SUCCESS
	if blogger.Id <= 0 {
		if err := blogger.Create(); err != nil {
			code = response.ERROR
		}
	} else {
		if err := blogger.UpdateInfo(); err != nil {
			code = response.ERROR
		}
	}

	response.CodeResponse(c, code)
}