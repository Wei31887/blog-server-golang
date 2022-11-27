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

