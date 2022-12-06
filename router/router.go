package router

import (
	"blog/server/admin"
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

	// Add middleware: cor
	router.Use(middleware.Cors())

	// Add middleware: log
	router.Use(middleware.Logger())

	fronEndRegister(router)
	AdminRegister(router)
	return router
}

func fronEndRegister(router *gin.Engine) {
	// Query blogger information
	router.POST("/blogger", api.FindBlogger)
	// Query total count of blog type
	router.POST("/blog/type", api.FindType)
	// Query blof list
	router.POST("/blog/list", api.BlogList)
	// Query content of blog
	router.POST("/blog/show", api.FindBlog)
	// Create new comment
	router.POST("/blog/comment", api.CreateComment)
}

func AdminRegister(router *gin.Engine) {
	router.POST("/login", admin.Login)
	router.POST("/logout", admin.Logout)
}