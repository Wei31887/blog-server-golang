package service

import (
	G "blog/server/global"
	"blog/server/model"
	"blog/server/utils"

	"gorm.io/gorm"
)

type BlogService struct{}

func (*BlogService) Create(blog *model.Blog) error {
	return G.GLOBAL_DB.Model(blog).Create(blog).Error
}

func (*BlogService) Update(blog *model.Blog) error {
	return G.GLOBAL_DB.Updates(blog).Error
}

func (*BlogService) Delete(blog *model.Blog) error {
	return G.GLOBAL_DB.Model(blog).Delete(blog).Error
}

func (*BlogService) FindOne(blog *model.Blog) (*model.Blog, error) {
	resBlog := &model.Blog{}
	err := G.GLOBAL_DB.Where("id = ?", blog.Id).First(resBlog).Error
	return resBlog, err
}

// FindNextBlog : query the next blog of given blog
func (*BlogService) FindNextBlogWithType(blog *model.Blog) (*model.Blog, error) {
	resBlog := &model.Blog{}
	err := G.GLOBAL_DB.
		Select("blog.id, blog.title, blog.type_id").
		Where("id > ?", blog.Id).
		First(resBlog).Error
	return resBlog, err
}

// FindPreviosBlog : query the previous blog of given blog
func (*BlogService) FindPrevBlogWithType(blog *model.Blog) (*model.Blog, error) {
	resBlog := &model.Blog{}
	err := G.GLOBAL_DB.
		Select("blog.id, blog.title, blog.type_id").
		Where("id < ?", blog.Id).
		Order("id desc").
		First(resBlog).Error
	return resBlog, err
}

// FindBlogWithTypeName : query the blog with type name by the given blog id
func (*BlogService) FindBlogWithTypeName(blog *model.Blog) (*model.Blog, error) {
	resBlog := &model.Blog{}
	err := G.GLOBAL_DB.Table("blog").
		Select("*, blog_type.name as type_name").
		Joins("left join blog_type on blog.type_id = blog_type.id").
		Where("blog.id = ?", blog.Id).
		Order("blog_type.sort asc").
		Find(resBlog).Error
	return resBlog, err
}

// FindList : query the blog list of the page
func (*BlogService) FindList(blog *model.Blog, page *utils.Page) ([]*model.Blog, error) {
	blogList := make([]*model.Blog, 0)
	curDB := G.GLOBAL_DB.Table("blog").
		Select("blog.id, title, type_id, add_time, update_time, click_hit, blog_type.name as type_name").
		Joins("left join blog_type on blog.type_id = blog_type.id")

	if blog.TypeId > 0 {
		curDB = curDB.Where("blog.type_id = ? ", blog.TypeId)
	}

	// Limit the maximum query number and offset
	err := curDB.
		Limit(page.Size).
		Offset(page.GetStartPage()).
		Order("add_time desc").
		Find(&blogList).Error
	return blogList, err
}

func (*BlogService) FindCountByTypeId(blog *model.Blog) (count int64, err error) {
	err = G.GLOBAL_DB.Model(&model.Blog{}).
		Where("id = ?", blog.TypeId).
		Count(&count).Error
	return
}

// Count the total page of the blog
func (*BlogService) Count() (count int64) {
	G.GLOBAL_DB.Model(&model.Blog{}).Count(&count)
	return
}

func (*BlogService) UpdataClick(blog *model.Blog) error {
	err := G.GLOBAL_DB.Model(blog).
		Update("click_hit", gorm.Expr("click_hit + 1")).Error
	return err
}

func (*BlogService) UpdateReplay(blog *model.Blog) error {
	err := G.GLOBAL_DB.Model(blog).
		Where("id = ? ", blog.Id).
		Update("replay_hit", gorm.Expr("replay_hit + ?", 1)).Error
	return err
}

// FindBlogComment : query the comment of the blog
func (*BlogService) FindBlogComment(blog *model.Blog) ([]*model.Comment, error) {
	comments := make([]*model.Comment, 0)
	err := G.GLOBAL_DB.Table("comment").
		Where("blog_id = ? and status = 0", blog.Id).
		Order("add_time desc").
		Find(&comments).Error
	return comments, err
}

// BlogListWithTag
func (*BlogService) BlogListWithTag(tagIdList []int, page *utils.Page) ([]*model.Blog, error) {
	blogList := make([]*model.Blog, 0)
	db := G.GLOBAL_DB.Model(&model.Blog{}).
		Select("blog.id, title, type_id, add_time, update_time, click_hit, blog_type.name as type_name, BLOG_TAG.TAG_ID AS TAG_ID").
		Joins("INNER JOIN BLOG_TAG ON BLOG.ID = BLOG_TAG.BLOG_ID").
		Joins("LEFT JOIN blog_type on BLOG.type_id = blog_type.id")

	for _, tag := range tagIdList {
		db = db.Where("tag_id = ?", tag)
	}

	err := db.Limit(page.Size).
		Offset(page.GetStartPage()).
		Order("BLOG.ID ASC").
		Find(&blogList).Error
	return blogList, err
}
