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
		response.CodeResponse(c, response.BADREQUEST)
		return
	}

	var comment service.Comment
	if page.Total, err = comment.Count(); err != nil {
		response.CodeResponse(c, response.ERROR)
		return	
	}

	result, err := comment.FindCommentList(page) 
	if err != nil {
		response.CodeResponse(c, response.ERROR)
		return	
	}
	res := response.Response{
		Data: result,
		Count: page.Total,
	}
	res.Json(c)
}

// CommentDelete : Query to delete the certain comment
func CommentDelete(c *gin.Context) {
	var comment service.Comment
	if err := c.BindJSON(&comment); err != nil {
		response.CodeResponse(c, response.BADREQUEST)
		return
	}

	err := comment.Delete()
	if err != nil {
		response.CodeResponse(c, response.ERROR)
		return
	}
	response.CodeResponse(c, response.SUCCESS)
}

func CommentStatus(c *gin.Context) {
	var comment service.Comment
	err := c.BindJSON(&comment)
	if err != nil {
		response.CodeResponse(c, response.BADREQUEST)
		return
	}

	err = comment.UpdateState()
	if err != nil {
		response.CodeResponse(c, response.ERROR)
		return	
	}
	response.CodeResponse(c, response.SUCCESS)
}