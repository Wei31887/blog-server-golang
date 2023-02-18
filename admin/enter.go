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
	blogService = service.Service.BloggerService
)

