package admin

import "blog/server/service"

type AdminApi struct {
	// TokenMaker token.Maker
	AdminBlogApi
	AdminBloggerApi
	AdminBlogTypeApi
	AdminCommentApi
	AdminBlogTagApi
	AdminTokenApi
}

var (
	bloggerService  = service.Service.BloggerService
	blogService     = service.Service.BlogService
	blogTypeService = service.Service.BlogTypeService
	commentService  = service.Service.CommentService
	tagService      = service.Service.TagService
	blogTagService  = service.Service.BlogTagService
	sessionService  = service.Service.SessionService
)
