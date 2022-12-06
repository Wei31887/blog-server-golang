package service

import (
	G "blog/server/global"
	"blog/server/model"
	"blog/server/utils"
)

type Comment model.Comment

// Gorm 約定 table name
func (Comment) TableName() string {
	return "comment"
}

func (commment *Comment) FindCommentList(page *utils.Page) {

}

// create comment
func (comment *Comment) Create() error {
	Db := G.GLOBAL_DB.Create(comment)
	return Db.Error
}

// update comment
func (comment *Comment) UpdateState() error {
	Db := G.GLOBAL_DB.Where("id = ?", comment.Id).Update("status", 1)
	return Db.Error
}

// delete comment
func (comment *Comment) Delete() error {
	Db := G.GLOBAL_DB.Model(comment).Delete(comment)
	return Db.Error
}