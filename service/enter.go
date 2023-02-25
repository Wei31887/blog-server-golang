package service

type ServiceGroup struct {
	BloggerService
	BlogTypeService
	BlogService
	CommentService
	TagService
	BlogTagService
	SessionService
}

var Service = &ServiceGroup{}
