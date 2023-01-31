package service

import (
	G "blog/server/global"
	"blog/server/model"
)

type Tag model.Tag
type BlogTag model.BlogTag

// Gorm 約定 table name
func (Tag) TableName() string {
	return "tag"
}
func (BlogTag) TableName() string {
	return "blog_tag"
}

/* ---- Tag ---- */
func (t *Tag) Create() error {
	db := G.GLOBAL_DB.Create(t)
	return db.Error
}

func (t *Tag) Update() error {
	db := G.GLOBAL_DB.Save(t)
	return db.Error
}

func (t *Tag) Delete() error {
	db := G.GLOBAL_DB.Delete(t)
	return db.Error
}

// TagList
func (t *Tag) TagList() ([]*Tag, error) {
	var tagList = make([]*Tag, 0)
	if db := G.GLOBAL_DB.Order("sort asc").Find(&tagList); db.Error != nil {
		return nil, db.Error
	}
	return tagList, nil
}



/* ---- BlogTag imtermidary table ---- */
func (bt *BlogTag) Create() error {
	db := G.GLOBAL_DB.Create(bt)
	return db.Error
}

func (bt *BlogTag) Update() error {
	db := G.GLOBAL_DB.Save(bt)
	return db.Error
}

func (bt *BlogTag) Delete() error {
	db := G.GLOBAL_DB.Delete(bt)
	return db.Error
}