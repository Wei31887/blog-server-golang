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

	// router for front end
	frontEndRegister(router)
	// router for admin
	AdminRegister(router)
	return router
}

func frontEndRegister(router *gin.Engine) {
	// Query blogger information
	router.POST("/blogger", api.FindBlogger)
	// Query total count of blog type
	router.POST("/blog/type", api.FindType)
	// Query blog list
	router.POST("/blog/list", api.BlogList)
	// Query content of blog
	router.POST("/blog/content", api.FindBlog)
	// Create new comment
	router.POST("/blog/comment", api.CreateComment)
	// Query Tag list 
	router.POST("/tag/list", api.TagList)
	// Query Blog List with Tag
	router.POST("/tag/blog", api.BlogListWithTag)
}

func AdminRegister(router *gin.Engine) {
	router.POST("/login", admin.Login)
	router.POST("/logout", admin.Logout)

	jwt := router.Group("/admin", middleware.JWT())

	{
		jwt.POST("/blogger/find", admin.FindBlogger)
		jwt.POST("/blogger/updatePassword", admin.UpdatePassword)
		jwt.POST("/blogger/updateInfo", admin.UpdateInfo)
		jwt.POST("/blog/type/list", admin.BlogTypeList)
		jwt.POST("/blog/type/save", admin.BlogTypeSave)
		jwt.POST("/blog/type/one", admin.BlogTypeOne)
		jwt.POST("/blog/type/delete", admin.BlogTypeDelete)
		jwt.POST("/blog/type/All", admin.BlogTypeAll)
		jwt.POST("/blog/save", admin.BlogSave)
		jwt.POST("/blog/list", admin.BlogList)
		jwt.POST("/blog/one", admin.BlogFindOne)
		jwt.POST("/blog/delete", admin.BlogDelete)
		jwt.POST("/comment/list", admin.CommentList)
		jwt.POST("/comment/updateStatus", admin.CommentStatus)
		jwt.POST("/comment/delete", admin.CommentDelete)
	}
}