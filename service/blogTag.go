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
func (tag *Tag) Create() error {
	db := G.GLOBAL_DB.Create(tag)
	return db.Error
}

func (tag *Tag) Update() error {
	db := G.GLOBAL_DB.Save(tag)
	return db.Error
}

func (tag *Tag) Delete() error {
	db := G.GLOBAL_DB.Delete(tag)
	return db.Error
}

// TagList : get the tag list with the count of each tag
func (tag *Tag) TagList() ([]*Tag, error) {
	var tagList = make([]*Tag, 0)
	db := G.GLOBAL_DB.Model(tag).
					Select("tag.id, tag.tag_name, tag.sort, count(blog_tag.blog_id) as count").
					Joins("left join blog_tag on tag.id=blog_tag.tag_id").
					Group("tag.id").
					Order("tag.id asc").Find(&tagList)
	if db.Error != nil {
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