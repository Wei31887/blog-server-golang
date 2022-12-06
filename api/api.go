package api

import (
	G "blog/server/global"
	"blog/server/model/response"
	"blog/server/service"
	"blog/server/utils"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

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
	result, err := blogType.FindTypeCount()
	
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
	json := make(map[string]interface{}, 0)
	err := c.ShouldBind(&json)
	if err != nil {
		res := response.Response{
			Code: response.INVALID_PARAMS,
			Msg: response.GetMsg(response.INVALID_PARAMS),
		}
		res.Json(c)
		return
	}

	// organize the query info
	blog := new(service.Blog)
	page, _ := strconv.Atoi(utils.ParseJsonString(json["page"]))
	size, _ := strconv.Atoi(utils.ParseJsonString(json["size"]))
	pageInfo := &utils.Page{
		Page: page,
		Size: size,
		Total: int(blog.Count()),
	}

	// get the type id if it exist
	typeId, err := strconv.Atoi(utils.ParseJsonString(json["type_id"]))
	if err == nil {
		blog.TypeId = typeId
	}
	// query the blog list
	results, err := blog.FindList(*pageInfo)
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
	prevBlog, _ := blog.FindPreviousBlog()
	// Query the next blog
	nextBlog, _ := blog.FindNextBlog()
	// Query the comments of the blog
	comments, _ := blog.FindBlogComment()

	resMap := make(map[string]interface{})
	resMap["last"] = prevBlog
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