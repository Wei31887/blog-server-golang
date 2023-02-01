package api

import (
	G "blog/server/global"
	"blog/server/model/response"
	"blog/server/service"
	"blog/server/utils"
	"time"

	"github.com/gin-gonic/gin"
)



type blogListInfo struct {
	TypeId int `json:"type_id"`
	Page int `json:"page"`
	Size int `json:"size"`
	Tags []TagID `json:"tags" binding:"dive"`
}
type TagID struct {
	ID int `json:"tag_id"`
}

// FindBlogger : request the blogger information
func FindBlogger(c *gin.Context) {
	var blogger service.Blogger
	result , err := blogger.FindIdFirst()
	if err != nil {
		res := response.Response{
			Code: response.ERROR,
			Msg: response.GetMsg(response.ERROR),
		}
		res.Json(c)
		return
	}

	result.Password = ""
	res := response.Response{
		Code: response.SUCCESS,
		Msg: response.GetMsg(response.SUCCESS),
		Data: result,
	}
	res.Json(c)
}

// FindType : request the amount of each type
func FindType(c *gin.Context) {
	var blogType service.BlogType
	result, err := blogType.FindAllTypeCount()
	
	if err != nil {
		res := response.Response{
			Code: response.ERROR,
			Msg: response.GetMsg(response.ERROR),
			Data: result,
		}
		res.Json(c)
		return
	}
	res := response.Response{
		Code: response.SUCCESS,
		Msg: response.GetMsg(response.SUCCESS),
		Data: result,
	}
	res.Json(c)
}

// BlogList : request the blog list of one page
func BlogList(c *gin.Context) {
	var requestInfo blogListInfo
	if err := c.BindJSON(&requestInfo); err != nil {
		response.ResponseWithCode(c, response.INVALID_PARAMS)
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
		Data: results,
		Count: pageInfo.Total,
	}
	res.Json(c)
}

// FindBlog : request the information of blog including comment, next blog, last page
func FindBlog(c *gin.Context) {
	var blog service.Blog
	err := c.BindJSON(&blog)
	if err != nil {
		res := response.Response {
			Code: response.INVALID_PARAMS,
			Msg: response.GetMsg(response.INVALID_PARAMS),
		}
		res.Json(c)
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
		Code: response.SUCCESS,
		Msg: response.GetMsg(response.SUCCESS),
		Data: resMap,
	}
	res.Json(c)
}

// CreateComment : api to create comment
func CreateComment(c *gin.Context) {
	var comment service.Comment
	err := c.BindJSON(&comment)
	if err != nil {
		res := response.Response{
			Code: response.INVALID_PARAMS,
			Msg: response.GetMsg(response.INVALID_PARAMS),
		}
		res.Json(c)
		return
	}

	// add ip and time to comment struct
	comment.Ip = c.ClientIP()
	comment.AddTime = time.Now().Format(G.DateFormat)

	err = comment.Create()
	if err != nil {
		res := response.Response {
			Code: response.ERROR,
			Msg: response.GetMsg(response.ERROR),
		}
		res.Json(c)
		return
	}

	// update blog
	blog := service.Blog{
		Id: comment.BlogId,
	}
	blog.UpdateReplay()

	res := response.Response {
		Code: response.SUCCESS,
		Msg: response.GetMsg(response.SUCCESS),
	}
	res.Json(c)
}

// TagList : get the list of tag
func TagList(c *gin.Context) {
	var tag service.Tag
	result, err := tag.TagList()
	if err != nil {
		response.ResponseWithCode(c, response.ERROR)
		return
	}

	response.SuccessWithData(c, result)
}

func BlogListWithTag(c *gin.Context) {
	var requestInfo blogListInfo
	if err := c.BindJSON(&requestInfo); err != nil {
		response.ResponseWithCode(c, response.INVALID_PARAMS)
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
	
	res, err := blog.BlogListWithTag(tagIdList, pageInfo)
	if err != nil {
		response.ResponseWithCode(c, response.ERROR)
		return
	}

	response.SuccessWithData(c, res)
}