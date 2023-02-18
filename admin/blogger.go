package admin

import (
	G "blog/server/global"
	"blog/server/model"
	"blog/server/model/response"
	"blog/server/token"
	"database/sql"

	"github.com/gin-gonic/gin"
)

type AdminBloggerApi struct {}

// Login
func (*AdminBloggerApi) Login(c *gin.Context) {
	blogger := &model.Blogger{}
	err := c.ShouldBindJSON(blogger)
	if err != nil {
		response.CodeResponse(c, response.BADREQUEST)
		return
	}

	resultBlogger, err := blogService.FindByName(blogger)
	if err != nil {
		if err == sql.ErrNoRows {
			response.MsgResponse(c, response.FORBIDDEN, "User not found")
			return
		}
		response.CodeResponse(c, response.ERROR)
		return
	}

	if blogger.Password != resultBlogger.Password {
		response.MsgResponse(c, response.FORBIDDEN, "Wrong password!")
		return
	}

	// create JWT token 
	maker := token.NewJWTMaker(G.GLOBAL_CONFIG.JWT.SigningKey)
	tokenStr, _, err := maker.CreateToken(blogger.Username, G.GLOBAL_CONFIG.JWT.ExpireTime)
	if err != nil {
		response.CodeResponse(c, response.ERROR)
		return
	}

	res := response.Response{
		Data: tokenStr,
	}
	res.Json(c)
}

// Logout
func (*AdminBloggerApi) Logout(c *gin.Context) {
	maker := token.NewJWTMaker(G.GLOBAL_CONFIG.JWT.SigningKey)
	if err := maker.JoinBlackList(c.GetHeader("token")); err != nil {
		response.CodeResponse(c, response.ERROR_AUTH_CHECK_TOKEN_IN_BLACK_LIST)
		return
	}
	response.CodeResponse(c, response.SUCCESS)
}

//
func (*AdminBloggerApi) FindBlogger(c *gin.Context) {
	queryBlogger, err := blogService.FindIdFirst()
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
func (*AdminBloggerApi) UpdatePassword(c *gin.Context) {
	blogger := &model.Blogger{}
	err := c.ShouldBindJSON(blogger)
	if err != nil {
		response.CodeResponse(c, response.BADREQUEST)
		return
	}

	code := response.SUCCESS
	if blogger.Id <= 0 {
		if err = blogService.Create(blogger); err != nil {
			code = response.ERROR
		}
	} else {
		if _, err = blogService.UpdateSecurityInfo(blogger); err != nil {
			code = response.ERROR
		}
	}

	response.CodeResponse(c, code)
}

func (*AdminBloggerApi) UpdateInfo(c *gin.Context) {
	blogger := &model.Blogger{}
	err := c.ShouldBindJSON(blogger)
	if err != nil {
		response.CodeResponse(c, response.BADREQUEST)
		return
	}

	code := response.SUCCESS
	if blogger.Id <= 0 {
		if err := blogService.Create(blogger); err != nil {
			code = response.ERROR
		}
	} else {
		if _, err := blogService.UpdateInfo(blogger); err != nil {
			code = response.ERROR
		}
	}

	response.CodeResponse(c, code)
}