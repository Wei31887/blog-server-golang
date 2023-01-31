package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Encapsulation of response message
// Define the format of response message
type Response struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
	Data interface{} `json:"data,omitempty"`
	Count int `json:"count,omitempty"`
}

// 
func (res *Response) Json (c *gin.Context) {
	c.JSON(http.StatusOK, res)
}

// success response with only message
func SuccessResponse(c *gin.Context) {
	res := Response {
		Code: SUCCESS,
		Msg: GetMsg(SUCCESS),
	}
	res.Json(c)
}

// success response with message and data
func SuccessWithData(c *gin.Context, data interface{}) {
	res := Response {
		Code: SUCCESS,
		Msg: GetMsg(SUCCESS),
		Data: data,
	}
	res.Json(c)
}

// success response with message and data
func SuccessWithDataCount(c *gin.Context, data interface{}, count int) {
	res := Response {
		Code: SUCCESS,
		Msg: GetMsg(SUCCESS),
		Data: data,
		Count: count,
	}
	res.Json(c)
}

// response with code and msg
func ResponseWithCode(c *gin.Context, code int) {
	res := Response {
		Code: code,
		Msg: GetMsg(code),
	}
	res.Json(c)
}

// response with code and msg
func ResponseCodeMsg(c *gin.Context, code int, msg string) {
	res := Response {
		Code: code,
		Msg: msg,
	}
	res.Json(c)
}
