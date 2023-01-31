package admin

import (
	"blog/server/model/response"
	"blog/server/service"
	"blog/server/utils"

	"github.com/gin-gonic/gin"
)


func CommentList(c *gin.Context) {
	var page utils.Page
	var err error
	if err = c.BindJSON(&page); err != nil {
		response.ResponseWithCode(c, response.INVALID_PARAMS)
		return
	}

	var comment service.Comment
	if page.Total, err = comment.Count(); err != nil {
		response.ResponseWithCode(c, response.ERROR)
		return	
	}

	result, err := comment.FindCommentList(page) 
	if err != nil {
		response.ResponseWithCode(c, response.ERROR)
		return	
	}
	response.SuccessWithDataCount(c, result, page.Total)
}

// CommentDelete : Query to delete the certain comment
func CommentDelete(c *gin.Context) {
	var comment service.Comment
	if err := c.BindJSON(&comment); err != nil {
		response.ResponseWithCode(c, response.INVALID_PARAMS)
		return
	}

	err := comment.Delete()
	if err != nil {
		response.ResponseWithCode(c, response.ERROR)
		return
	}
	response.SuccessResponse(c)
}

func CommentStatus(c *gin.Context) {
	var comment service.Comment
	err := c.BindJSON(&comment)
	if err != nil {
		response.ResponseWithCode(c, response.INVALID_PARAMS)
		return
	}

	err = comment.UpdateState()
	if err != nil {
		response.ResponseWithCode(c, response.ERROR)
		return	
	}
	response.SuccessResponse(c)
}