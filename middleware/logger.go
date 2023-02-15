package middleware

import (
	"blog/server/global"
	"io"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type LogLayout struct {
	Code          int
	CostTime      time.Duration
	ClientIp      string
	RequestMethod string
	RequestURI    string
	Header        map[string][]string
	Protocl       string
	Error         string
	Body          string
}

// Logger : midlleware logger
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// start time
		startTime := time.Now()
		c.Next()
		// during time of request
		latencyTime := time.Since(startTime)
		// body of request
		body, _ := io.ReadAll(c.Request.Body)

		layout := LogLayout{
			Code:          c.Writer.Status(),
			CostTime:      latencyTime,
			ClientIp:      c.ClientIP(),
			RequestMethod: c.Request.Method,
			RequestURI:    c.Request.RequestURI,
			Header:        c.Request.Header,
			Protocl:       c.Request.Proto,
			Body:          string(body),
		}

		// write the log
		logger := global.GLOBAL_LOG
		logger.Info(
			"Request info:",
			zap.Int("status-code", layout.Code),
			zap.Duration("cost-time", layout.CostTime),
			zap.String("client-ip", layout.ClientIp),
			zap.String("request-method", layout.RequestMethod),
			zap.String("request-uri", layout.RequestURI),
			// zap.String("header", layout.Header),
			zap.String("protocl", layout.Protocl),

		)
	}
}
