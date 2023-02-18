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
	blogService = service.Service.BloggerService
	blogTypeService = service.Service.BlogTypeService
	commentService = service.Service.CommentService
)