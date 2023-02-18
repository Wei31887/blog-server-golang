package router

import (
	"blog/server/middleware"

	"github.com/gin-gonic/gin"
)


func InitRouter() *gin.Engine {
	router := gin.New()
	// gin.Recovery: 
	// Recovery returns a middleware that recovers from any panics and writes a 500 if there was one.
	router.Use(gin.Recovery())
	router.Static("/static", "static")

	// Add middleware
	router.Use(middleware.Cors())
	router.Use(middleware.Logger())

	ApiGroupfrontEndRegister(router)
	ApiGroupAdminRegister(router)

	return router
}

func ApiGroupfrontEndRegister(router *gin.Engine) {
	router.POST("/blogger", apiGroup.FindBlogger)
	router.POST("/blog/type", apiGroup.FindType)
	router.POST("/blog/list", apiGroup.BlogList)
	router.POST("/blog/content", apiGroup.FindBlog)
	router.POST("/blog/comment", apiGroup.CreateComment)
	router.POST("/tag/list", apiGroup.TagList)
	router.POST("/tag/blog", apiGroup.BlogListWithTag)
}

func ApiGroupAdminRegister(router *gin.Engine) {
	router.POST("/login", adminApiGroup.Login)
	
	authRouter := router.Group("", middleware.JWT())
	authRouter.POST("/logout", adminApiGroup.Logout)
	
	jwt := router.Group("/admin", middleware.JWT())
	{	
		jwt.POST("/blogger/find", adminApiGroup.FindBlogger)
		jwt.POST("/blogger/updatePassword", adminApiGroup.UpdatePassword)
		jwt.POST("/blogger/updateInfo", adminApiGroup.UpdateInfo)
		jwt.POST("/blog/type/list", adminApiGroup.BlogTypeList)
		jwt.POST("/blog/type/save", adminApiGroup.BlogTypeSave)
		jwt.POST("/blog/type/one", adminApiGroup.BlogTypeOne)
		jwt.POST("/blog/type/delete", adminApiGroup.BlogTypeDelete)
		jwt.POST("/blog/type/All", adminApiGroup.BlogTypeAll)
		jwt.POST("/blog/save", adminApiGroup.BlogSave)
		jwt.POST("/blog/list", adminApiGroup.BlogList)
		jwt.POST("/blog/one", adminApiGroup.BlogFindOne)
		jwt.POST("/blog/delete", adminApiGroup.BlogDelete)
		jwt.POST("/comment/list", adminApiGroup.CommentList)
		jwt.POST("/comment/updateStatus", adminApiGroup.CommentStatus)
		jwt.POST("/comment/delete", adminApiGroup.CommentDelete)
	}
}