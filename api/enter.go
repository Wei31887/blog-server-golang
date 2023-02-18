package api

import "blog/server/service"

type Api struct {
	// TokenMaker token.Maker
	BlogApi
	BloggerApi
	BlogTypeApi
	CommentApi
	TagApi
}

var (
	bloggerService  = service.Service.BloggerService
	blogService     = service.Service.BlogService
	blogTypeService = service.Service.BlogTypeService
	commentService  = service.Service.CommentService
	tagService      = service.Service.TagService
)
