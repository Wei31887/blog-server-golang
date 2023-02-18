package service

import (
	G "blog/server/global"
	"blog/server/model"
)

type Tag model.Tag
type BlogTag model.BlogTag

type TagService struct{}
type BlogTagService struct{}

/* ---- Tag ---- */

func (t *TagService) Create(tag *model.Tag) error {
	return G.GLOBAL_DB.Create(tag).Error
}

func (t *TagService) Update(tag *model.Tag) error {
	return G.GLOBAL_DB.Updates(tag).Error
}

func (t *TagService) Delete(tag *model.Tag) error {
	return G.GLOBAL_DB.Delete(tag).Error
}

// TagList : get the tag list with the count of each tag
func (t *TagService) TagList() ([]*model.Tag, error) {
	var tagList = make([]*model.Tag, 0)
	err := G.GLOBAL_DB.Model(&model.Tag{}).
		Select("tag.id, tag.tag_name, tag.sort, count(blog_tag.blog_id) as count").
		Joins("left join blog_tag on tag.id=blog_tag.tag_id").
		Group("tag.id").
		Order("tag.id asc").
		Find(&tagList).Error
	return tagList, err
}

/* ---- BlogTag imtermidary table ---- */
func (bt *BlogTagService) Create(blogTag *model.BlogTag) error {
	db := G.GLOBAL_DB.Create(blogTag)
	return db.Error
}

func (bt *BlogTagService) Update(blogTag *model.BlogTag) error {
	db := G.GLOBAL_DB.Save(blogTag)
	return db.Error
}

func (bt *BlogTagService) Delete(blogTag *model.BlogTag) error {
	db := G.GLOBAL_DB.Delete(blogTag)
	return db.Error
}
