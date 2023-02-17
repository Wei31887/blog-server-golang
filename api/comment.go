package api

import (
	G "blog/server/global"
	"blog/server/model/response"
	"blog/server/service"
	"time"

	"github.com/gin-gonic/gin"
)

// CreateComment : api to create comment
func CreateComment(c *gin.Context) {
	var comment service.Comment
	err := c.BindJSON(&comment)
	if err != nil {
		response.CodeResponse(c, response.BADREQUEST)
		return
	}

	// add ip and time to comment struct
	comment.Ip = c.ClientIP()
	comment.AddTime = time.Now().Format(G.DateFormat)

	err = comment.Create()
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