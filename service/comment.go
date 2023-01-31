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

func (commment *Comment) FindCommentList(page utils.Page) ([]*Comment, error) {
	commentLists := make([]*Comment, 0)
	db := G.GLOBAL_DB.Model(commment).Limit(page.Size).Offset(page.GetStartPage()).Find(&commentLists)
	if (db.Error != nil ) {
		return nil, db.Error
	} 
	return commentLists, nil
}

func (commment *Comment) Count() (int, error) {
	var count int64
	db := G.GLOBAL_DB.Model(commment).Count(&count)
	if (db.Error != nil ) {
		return 0, db.Error
	} 
	return int(count), nil
}

// create comment
func (comment *Comment) Create() error {
	Db := G.GLOBAL_DB.Create(comment)
	return Db.Error
}

// update comment
func (comment *Comment) UpdateState() error {
	Db := G.GLOBAL_DB.Model(comment).Where("id = ?", comment.Id).Update("status", comment.Status)
	return Db.Error
}

// delete comment
func (comment *Comment) Delete() error {
	Db := G.GLOBAL_DB.Model(comment).Delete(comment)
	return Db.Error
}