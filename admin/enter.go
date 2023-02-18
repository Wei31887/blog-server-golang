package admin

import "blog/server/service"

type AdminApi struct {
	// TokenMaker token.Maker
	AdminBlogApi
	AdminBloggerApi
	AdminBlogTypeApi
	AdminCommentApi
}

var (
	bloggerService  = service.Service.BloggerService
	blogService     = service.Service.BlogService
	blogTypeService = service.Service.BlogTypeService
	commentService  = service.Service.CommentService
)
