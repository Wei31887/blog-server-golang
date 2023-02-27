package server

import (
	"blog/server/initialize/config"
	"blog/server/middleware"

	"github.com/gin-gonic/gin"
)

type HTTPServer struct {
	Router     *gin.Engine
	Config     *config.Config
}

func NewHTTPServer(config *config.Config) *HTTPServer {
	
	server := &HTTPServer{
		Config: config,
	}
	server.setupRouter()
	return server
}

func (server *HTTPServer) setupRouter() {
	router := gin.New()
	router.Use()
	router.Use(gin.Recovery())
	router.Static("/static", "static")

	// Add middleware
	router.Use(middleware.Cors())
	router.Use(middleware.Logger())

	ApiGroupfrontEndRegister(router)
	ApiGroupAdminRegister(router)
	server.Router = router
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
	router.POST("/token/refresh", adminApiGroup.RefreshToken)

	jwt := router.Group("/admin", middleware.JWT())
	{
		jwt.POST("/logout", adminApiGroup.Logout)
		// blogger
		jwt.POST("/blogger/find", adminApiGroup.FindBlogger)
		jwt.POST("/blogger/updatePassword", adminApiGroup.UpdatePassword)
		jwt.POST("/blogger/updateInfo", adminApiGroup.UpdateInfo)

		// type
		jwt.POST("/blog/type/list", adminApiGroup.BlogTypeList)
		jwt.POST("/blog/type/save", adminApiGroup.BlogTypeSave)
		jwt.POST("/blog/type/one", adminApiGroup.BlogTypeOne)
		jwt.POST("/blog/type/delete", adminApiGroup.BlogTypeDelete)
		jwt.POST("/blog/type/all", adminApiGroup.BlogTypeAll)

		// tag
		jwt.POST("/blog/tag/all", adminApiGroup.BlogTagAll)
		jwt.POST("/blog/tag/list", adminApiGroup.BlogTagList)
		jwt.POST("/blog/tag/save", adminApiGroup.BlogTagSave)
		jwt.POST("/blog/tag/delete", adminApiGroup.BlogTagDelete)
		// jwt.POST("/blog/type/one", adminApiGroup.BlogTypeOne)

		// blog
		jwt.POST("/blog/save", adminApiGroup.BlogSave)
		jwt.POST("/blog/list", adminApiGroup.BlogList)
		jwt.POST("/blog/one", adminApiGroup.BlogFindOne)
		jwt.POST("/blog/delete", adminApiGroup.BlogDelete)

		// comment
		jwt.POST("/comment/list", adminApiGroup.CommentList)
		jwt.POST("/comment/updateStatus", adminApiGroup.CommentStatus)
		jwt.POST("/comment/delete", adminApiGroup.CommentDelete)
	}
}

func (server *HTTPServer) RunServer(address string) error {
	return server.Router.Run(address)
}
