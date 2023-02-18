package service

import (
	G "blog/server/global"
	"blog/server/model"
	"blog/server/utils"
)

// type Comment model.Comment
type CommentService struct {}

// var commentService CommentService

func (*CommentService) FindCommentList(page utils.Page) ([]*model.Comment, error) {
	commentLists := make([]*model.Comment, 0)
	err := G.GLOBAL_DB.Model(&model.Comment{}).
				Limit(page.Size).
				Offset(page.GetStartPage()).
				Find(&commentLists).Error
	return commentLists, err
}

func (*CommentService) Count() (int, error) {
	var count int64
	err := G.GLOBAL_DB.Model(&model.Comment{}).Count(&count).Error
	return int(count), err
}

// create comment
func (*CommentService) Create(comment *model.Comment) error {
	return G.GLOBAL_DB.Create(comment).Error
	
}

// update comment
func (*CommentService) UpdateState(comment *model.Comment) error {
	return G.GLOBAL_DB.Model(comment).Where("id = ?", comment.Id).Update("status", comment.Status).Error
}

// delete comment
func (*CommentService) Delete(comment *model.Comment) error {
	return G.GLOBAL_DB.Model(comment).Delete(comment).Error
}