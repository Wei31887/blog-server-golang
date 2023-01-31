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
		res := response.Response{
			Code: response.INVALID_PARAMS,
			Msg: response.GetMsg(response.INVALID_PARAMS),
		}
		res.Json(c)
		return
	}

	queryBloger, _ := blogger.FindByName()
	if queryBloger == nil {
		res := response.Response{
			Code: response.NOTFOUND,
			Msg: "User not found!",
		}
		res.Json(c)
		return
	}

	if blogger.Password != queryBloger.Password {
		res := response.Response{
			Code: response.INVALID_PARAMS,
			Msg: "Wrong password!",
		}
		res.Json(c)
		return
	}

	// create JWT token 
	j := utils.NewJWT()
	jwtToken, err := j.GenerateToken(blogger.Username)
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
		Data: jwtToken,
	}
	res.Json(c)
}

// Logout
func Logout(c *gin.Context) {
	// delete the jwt token
	// http.SetCookie(c.Writer, &http.Cookie{
	// 	Name:     "token",
	// 	MaxAge:   -1,
	// })
	res := response.Response{
		Code: response.SUCCESS,
		Msg: response.GetMsg(response.SUCCESS),
	}
	res.Json(c)
}

//
func FindBlogger(c *gin.Context) {
	var blogger service.Blogger
	queryBlogger, err := blogger.FindIdFirst()
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
		Data: queryBlogger,
	}
	res.Json(c)
}

// Update blogger password
func UpdatePassword(c *gin.Context) {
	var blogger service.Blogger
	err := c.BindJSON(&blogger)
	if err != nil {
		res := response.Response{
			Code: response.INVALID_PARAMS,
			Msg: response.GetMsg(response.INVALID_PARAMS),
		}
		res.Json(c)
		return
	}

	code := response.SUCCESS
	if blogger.Id <= 0 {
		if _, err = blogger.Create(); err != nil {
			code = response.ERROR
		}
	} else {
		if err = blogger.UpdatePassword(); err != nil {
			code = response.ERROR
		}
	}

	res := response.Response{
		Code: code,
		Msg: response.GetMsg(code),
	}
	res.Json(c)
}

func UpdateInfo(c *gin.Context) {
	var blogger service.Blogger
	err := c.BindJSON(&blogger)
	if err != nil {
		res := response.Response{
			Code: response.INVALID_PARAMS,
			Msg: response.GetMsg(response.INVALID_PARAMS),
		}
		res.Json(c)
		return
	}

	code := response.SUCCESS
	if blogger.Id <= 0 {
		if _, err := blogger.Create(); err != nil {
			code = response.ERROR
		}
	} else {
		if err := blogger.UpdateInfo(); err != nil {
			code = response.ERROR
		}
	}

	res := response.Response {
		Code: code,
		Msg: response.GetMsg(code),
	}
	res.Json(c)
}