package api

import (
	"blog/server/model"
	"blog/server/model/response"
	"blog/server/utils"

	"github.com/gin-gonic/gin"
)

type BlogApi struct{}

// FindBlog : request the information of blog including comment, next blog, last page
func (*BlogApi) FindBlog(c *gin.Context) {
	blog := &model.Blog{}
	err := c.ShouldBindJSON(&blog)
	if err != nil {
		response.CodeResponse(c, response.BADREQUEST)
		return
	}

	// Update the click hit
	blogService.UpdataClick(blog)

	resultBlog, _ := blogService.FindOne(blog)

	prevBlog, _ := blogService.FindPrevBlogWithType(blog)

	nextBlog, _ := blogService.FindNextBlogWithType(blog)
	
	comments, _ := blogService.FindBlogComment(blog)

	tags, _ := blogTagService.FindBlogTag(blog.Id)
	// if err != sql.ErrNoRows {
	// 	response.CodeResponse(c, response.ERROR)
    //     return
	// }

	resMap := make(map[string]interface{})
	resMap["prev"] = prevBlog
	resMap["next"] = nextBlog
	resMap["blog"] = resultBlog
	resMap["tags"] = tags
	resMap["comments"] = comments

	res := response.Response{
		Data: resMap,
	}
	res.Json(c)
}

type blogListRequest struct {
	TypeId int `json:"type_id"`
	Page   int `json:"page"`
	Size   int `json:"size"`
}

// BlogList : request the blog list of one page
func (*BlogApi) BlogList(c *gin.Context) {
	var requestInfo blogListRequest
	if err := c.ShouldBindJSON(&requestInfo); err != nil {
		response.CodeResponse(c, response.BADREQUEST)
		return
	}

	// organize
	blog := model.Blog{}
	pageInfo := utils.Page{
		Page:  requestInfo.Page,
		Size:  requestInfo.Size,
		Total: int(blogService.Count()),
	}

	// get the type id if it exist
	if requestInfo.TypeId != 0 {
		blog.TypeId = requestInfo.TypeId
	}

	// query the blog list
	results, err := blogService.FindList(&blog, &pageInfo)
	if err != nil {
		response.CodeResponse(c, response.ERROR)
		return
	}

	res := response.Response{
		Data:  results,
		Count: pageInfo.Total,
	}
	res.Json(c)
}

type blogListWithTagRequest struct {
	Page int     `json:"page"`
	Size int     `json:"size"`
	Tags []TagID `json:"tags" binding:"dive"`
}
type TagID struct {
	ID int `json:"tag_id"`
}

func (*BlogApi) BlogListWithTag(c *gin.Context) {
	var requestInfo blogListWithTagRequest
	if err := c.ShouldBindJSON(&requestInfo); err != nil {
		response.CodeResponse(c, response.BADREQUEST)
		return
	}

	// organize the query info
	tagIdList := make([]int, 0)
	for _, item := range requestInfo.Tags {
		tagIdList = append(tagIdList, item.ID)
	}
	pageInfo := &utils.Page{
		Page:  requestInfo.Page,
		Size:  requestInfo.Size,
		Total: int(blogService.Count()),
	}

	results, err := blogService.BlogListWithTag(tagIdList, pageInfo)
	if err != nil {
		response.CodeResponse(c, response.ERROR)
		return
	}

	res := response.Response{
		Data: results,
	}
	res.Json(c)
}
