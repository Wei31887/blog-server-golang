package router

import (
	"blog/server/api"
	"blog/server/middleware"

	"github.com/gin-gonic/gin"
)


func InitRouter() *gin.Engine {
	router := gin.New()
	// gin.Recovery: 
	// Recovery returns a middleware that recovers from any panics and writes a 500 if there was one.
	router.Use(gin.Recovery())
	router.Static("/static", "static")

	// cor
	router.Use(middleware.Cors())

	// log
	router.Use(middleware.Logger())

	register(router)
	return router
}

func register(router *gin.Engine) {
	// Query blogger information
	router.POST("/blogger", api.FindBlogger)
	// Query total count of blog type
	router.POST("/blog/type", api.FindType)
	// Query blof list
	router.POST("/blog/list", api.BlogList)
	// Query content of blog
	router.POST("/blog/show", api.FindBlog)

}