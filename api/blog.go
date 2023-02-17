package api

import (
	"blog/server/model/response"
	"blog/server/service"
	"blog/server/utils"

	"github.com/gin-gonic/gin"
)

// FindBlog : request the information of blog including comment, next blog, last page
func FindBlog(c *gin.Context) {
	var blog service.Blog
	err := c.ShouldBindJSON(&blog)
	if err != nil {
		response.CodeResponse(c, response.BADREQUEST)
		return
	}

	// Update the click hit
	blog.UpdataClick()
	// Query the blog with type name by the given blog id
	resultBlog, _ := blog.FindBlogWithTypeName()
	// Query the previous blog
	prevBlog, _ := blog.FindPrevBlogWithType()
	// // Query the next blog
	nextBlog, _ := blog.FindNextBlogWithType()
	// Query the comments of the blog
	comments, _ := blog.FindBlogComment()

	resMap := make(map[string]interface{})
	resMap["prev"] = prevBlog
	resMap["next"] = nextBlog
	resMap["blog"] = resultBlog
	resMap["comments"] = comments

	res := response.Response{
		Data: resMap,
	}
	res.Json(c)
}


type blogListRequest struct {
	TypeId int `json:"type_id"`
	Page int `json:"page"`
	Size int `json:"size"`
}

// BlogList : request the blog list of one page
func BlogList(c *gin.Context) {
	var requestInfo blogListRequest
	if err := c.BindJSON(&requestInfo); err != nil {
		response.CodeResponse(c, response.BADREQUEST)
		return
	}

	// organize
	blog := new(service.Blog)
	pageInfo := &utils.Page{
		Page: requestInfo.Page,
		Size: requestInfo.Size,
		Total: int(blog.Count()),
	}

	// get the type id if it exist
	if requestInfo.TypeId != 0 {
		blog.TypeId = requestInfo.TypeId
	}

	// query the blog list
	results, err := blog.FindList(pageInfo)
	if err != nil {
		response.CodeResponse(c, response.ERROR)
		return
	}

	res := response.Response{
		Data: results,
		Count: pageInfo.Total,
	}
	res.Json(c)
}

type blogListWithTagRequest struct {
	Page int `json:"page"`
	Size int `json:"size"`
	Tags []TagID `json:"tags" binding:"dive"`
}
type TagID struct {
	ID int `json:"tag_id"`
}

func BlogListWithTag(c *gin.Context) {
	var requestInfo blogListWithTagRequest
	if err := c.BindJSON(&requestInfo); err != nil {
		response.CodeResponse(c, response.BADREQUEST)
		return
	}

	// organize the query info
	blog := new(service.Blog)
	tagIdList := make([]int, 0)
	for _, item := range requestInfo.Tags {
		tagIdList = append(tagIdList, item.ID)
	}
	pageInfo := &utils.Page{
		Page: requestInfo.Page,
		Size: requestInfo.Size,
		Total: int(blog.Count()),
	}
	
	results, err := blog.BlogListWithTag(tagIdList, pageInfo)
	if err != nil {
		response.CodeResponse(c, response.ERROR)
		return
	}

	res := response.Response{
        Data: results,
	}
	res.Json(c)
}