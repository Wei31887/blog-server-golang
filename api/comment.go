package api

import (
	"blog/server/model"
	"blog/server/model/response"
	"blog/server/service"
	"time"

	"github.com/gin-gonic/gin"
)

type CommentApi struct {}

// CreateComment : api to create comment
func (*CommentApi) CreateComment(c *gin.Context) {
	comment := model.Comment{}
	err := c.ShouldBindJSON(&comment)
	if err != nil {
		response.CodeResponse(c, response.BADREQUEST)
		return
	}

	// add ip and time to comment struct
	comment.Ip = c.ClientIP()
	comment.AddTime = time.Now()

	err = commentService.Create(&comment)
	if err != nil {
		response.CodeResponse(c, response.ERROR)
		return
	}

	// update blog
	blog := service.Blog{
		Id: comment.BlogId,
	}
	blog.UpdateReplay()

	response.CodeResponse(c, response.SUCCESS)
}