package service


type ServiceGroup struct {
	BloggerService
	BlogTypeService
	CommentService
}

var Service = &ServiceGroup{}