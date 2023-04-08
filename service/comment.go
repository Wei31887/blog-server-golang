package service

import (
	"blog/server/initialize/global"
	"blog/server/model"
	adminresponse "blog/server/model/response/admin"
	"blog/server/utils"
)

// type Comment model.Comment
type CommentService struct{}

// var commentService CommentService

func (*CommentService) FindCommentList(page utils.Page) ([]*adminresponse.CommentListResponse, error) {
	commentLists := make([]*adminresponse.CommentListResponse, 0)
	err := global.GLOBAL_DB.Table("comment").
		Select("comment.add_time, comment.blog_id, blog.title, comment.content, comment.id, comment.ip, comment.nick_name, comment.status").
		Joins("left join blog on comment.blog_id = blog.id").
		Limit(page.Size).
		Offset(page.GetStartPage()).
		Find(&commentLists).Error
	return commentLists, err
}

func (*CommentService) Count() (int, error) {
	var count int64
	err := global.GLOBAL_DB.Model(&model.Comment{}).Count(&count).Error
	return int(count), err
}

// create comment
func (*CommentService) Create(comment *model.Comment) error {
	return global.GLOBAL_DB.Create(comment).Error

}

// update comment
func (*CommentService) UpdateState(comment *model.Comment) error {
	return global.GLOBAL_DB.Model(comment).Where("id = ?", comment.Id).Update("status", comment.Status).Error
}

// delete comment
func (*CommentService) Delete(comment *model.Comment) error {
	return global.GLOBAL_DB.Model(comment).Delete(comment).Error
}
