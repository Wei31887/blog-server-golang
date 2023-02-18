package admin

import (
	"blog/server/model"
	"blog/server/model/response"
	"blog/server/utils"

	"github.com/gin-gonic/gin"
)

type AdminCommentApi struct{}

func (*AdminCommentApi) CommentList(c *gin.Context) {
	var page utils.Page
	var err error
	if err = c.ShouldBindJSON(&page); err != nil {
		response.CodeResponse(c, response.BADREQUEST)
		return
	}

	if page.Total, err = commentService.Count(); err != nil {
		response.CodeResponse(c, response.ERROR)
		return
	}

	result, err := commentService.FindCommentList(page)
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

// CommentDelete : Query to delete the certain comment
func (*AdminCommentApi) CommentDelete(c *gin.Context) {
	comment := model.Comment{}
	if err := c.ShouldBindJSON(&comment); err != nil {
		response.CodeResponse(c, response.BADREQUEST)
		return
	}

	err := commentService.Delete(&comment)
	if err != nil {
		response.CodeResponse(c, response.ERROR)
		return
	}
	response.CodeResponse(c, response.SUCCESS)
}

func (*AdminCommentApi) CommentStatus(c *gin.Context) {
	comment := model.Comment{}
	err := c.ShouldBindJSON(&comment)
	if err != nil {
		response.CodeResponse(c, response.BADREQUEST)
		return
	}

	err = commentService.UpdateState(&comment)
	if err != nil {
		response.CodeResponse(c, response.ERROR)
		return
	}
	response.CodeResponse(c, response.SUCCESS)
}
