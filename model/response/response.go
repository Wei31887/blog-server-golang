package response

import (
	"github.com/gin-gonic/gin"
)

// Encapsulation of response message
// Define the format of response message
type Response struct {
	Data interface{} `json:"data,omitempty"`
	Count int `json:"count,omitempty"`
}

func (res *Response) Json(c *gin.Context) {
	c.JSON(SUCCESS, res)
}
 

// errResponse
func MsgResponse(c *gin.Context, code int, msg string) {
	c.JSON(code, gin.H {
		"msg": msg,
	})
}

// MsgResponse
func CodeResponse(c *gin.Context, code int) {
	c.JSON(code, gin.H {
		"msg": GetMsg(code),
	})
}

// ErrorResponse
func ErrResponseJson(err error) gin.H {
    return gin.H{
		"err": err.Error(),
	}
}