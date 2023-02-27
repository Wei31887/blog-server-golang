package service

import (
	"blog/server/initialize/global"
	"blog/server/model"
	"blog/server/utils"
)

// type tag model.Tag
// type blogTag model.BlogTag

type TagService struct{}
type BlogTagService struct{}

/* ---- Tag ---- */

func (t *TagService) Create(tag *model.Tag) error {
	return global.GLOBAL_DB.Create(tag).Error
}

func (t *TagService) Update(tag *model.Tag) error {
	return global.GLOBAL_DB.Updates(tag).Error
}

func (t *TagService) Delete(tag *model.Tag) error {
	return global.GLOBAL_DB.Delete(tag).Error
}

func (t *TagService) Count() (int, error) {
	var count int64
	err := global.GLOBAL_DB.Model(&model.Tag{}).Count(&count).Error
	return int(count), err
}

func (t *TagService) ListPage(page utils.Page) ([]*model.Tag, error) {
	blogTypes := make([]*model.Tag, 0)
	err := global.GLOBAL_DB.Model(&model.Tag{}).
		Limit(page.Size).
		Offset(page.GetStartPage()).
		Order("sort asc").
		Find(&blogTypes).Error
	return blogTypes, err
}

// TagList : get the tag list with the count of each tag
func (t *TagService) ListAll() ([]*model.Tag, error) {
	var tagList = make([]*model.Tag, 0)
	err := global.GLOBAL_DB.
		Select("tag.id, tag.tag_name, tag.sort, count(blog_tag.blog_id) AS count").
		Joins("inner join blog_tag on tag.id = blog_tag.tag_id").
		Group("tag.id").
		Order("tag.id asc").
		Find(&tagList).Error
	return tagList, err
}

// TagAll : get all of tag 
func (t *TagService) TagAll() ([]*model.Tag, error) {
	var tagList = make([]*model.Tag, 0)
	err := global.GLOBAL_DB.Model(&model.Tag{}).
		Select("tag.id, tag.tag_name, tag.sort").
		Order("tag.id asc").
		Find(&tagList).Error
	return tagList, err
}

/* ---- BlogTag imtermidary table ---- */
func (*BlogTagService) Create(blogTag *model.BlogTag) error {
	db := global.GLOBAL_DB.Create(blogTag)
	return db.Error
}

func (*BlogTagService) Update(blogTag *model.BlogTag) error {
	db := global.GLOBAL_DB.Where("id = ?", blogTag.Id).Updates(blogTag)
	return db.Error
}

func (*BlogTagService) Delete(blogTag *model.BlogTag) error {
	db := global.GLOBAL_DB.Delete(blogTag)
	return db.Error
}

func (*BlogTagService) OneBlogTag(blogId int) ([]model.BlogTag, error) {
	var tagList = make([]model.BlogTag, 0)
	err := global.GLOBAL_DB.Model(&model.BlogTag{}).
		Select("blog_tag.id, blog_tag.blog_id, blog_tag.tag_id").
		Where("blog_tag.blog_id = ?", blogId).
		Order("tag_id asc").
		Find(&tagList).Error
	return tagList, err
}

func (*BlogTagService) FindBlogTag(blogId int) ([]*model.Tag, error) {
	var tagList = make([]*model.Tag, 0)
    err := global.GLOBAL_DB.Model(&model.BlogTag{}).
        Select("tag.id, tag.tag_name, tag.sort").
		Joins("inner join tag on tag.id = blog_tag.tag_id").
        Where("blog_tag.blog_id =?", blogId).
        Order("tag.sort asc").
        Find(&tagList).Error
    return tagList, err
}