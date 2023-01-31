package service

import (
	G "blog/server/global"
	"blog/server/model"
	"blog/server/utils"

	"gorm.io/gorm"
)

type Blog model.Blog

// Gorm 約定 table name
func (Blog) TableName() string {
	return "blog"
}

func (blog *Blog) Create() error {
	db := G.GLOBAL_DB.Model(blog).Create(blog)
	return db.Error
}

func (blog *Blog) Update() error {
	db := G.GLOBAL_DB.Updates(blog)
	return db.Error
}

func (blog *Blog) Delete() error {
	db := G.GLOBAL_DB.Model(blog).Delete(blog)
	return db.Error
}

func (blog *Blog) FindOne() (*Blog, error) {
	resBlog := new(Blog)
	db := G.GLOBAL_DB.Where("id = ?", blog.Id).First(resBlog)
	if db.Error != nil {
		return nil, db.Error
	}
	return resBlog, nil
}


// FindNextBlog : query the next blog of given blog
func (blog *Blog) FindNextBlogWithType() (*Blog, error) {
	resBlog := new(Blog)
	db := G.GLOBAL_DB.Select("blog.id, blog.title, blog.typeid").Where("id > ?", blog.Id).First(resBlog)
	if db.Error != nil {
		return nil, db.Error
	}
	return resBlog, nil
}

// FindPreviosBlog : query the previous blog of given blog
func (blog *Blog) FindPrevBlogWithType() (*Blog, error) {
	resBlog := new(Blog)
	db := G.GLOBAL_DB.Select("blog.id, blog.title, blog.typeid").Where("id < ?", blog.Id).Order("id desc").First(resBlog)
	if db.Error != nil {
		return nil, db.Error
	}
	return resBlog, nil
}

// FindBlogWithTypeName : query the blog with type name by the given blog id
func (blog *Blog) FindBlogWithTypeName() (*Blog, error) {
	resBlog := new(Blog)
	db := G.GLOBAL_DB.Table("blog").Select("*, blog_type.name as type_name").
					Joins("left join blog_type on blog.typeid = blog_type.id").
					Where("blog.id = ?", blog.Id).Order("blog_type.sort asc").
					Find(resBlog)

	if db.Error != nil {
		return nil, db.Error
	}
	return resBlog, nil
}


// FindList : query the blog list of the page
func (blog *Blog) FindList(page *utils.Page) ([]*Blog, error) {
	blogList := make([]*Blog, 0)
	curDB := G.GLOBAL_DB.Table("blog").Select("blog.id, title, typeid, add_time, update_time, click_hit, blog_type.name as type_name").
				Joins("left join blog_type on blog.typeid = blog_type.id")

	if blog.TypeId > 0 {
		curDB = curDB.Where("blog.typeid = ? ", blog.TypeId)
	}
	
	// Limit the maximum query number and offset
	result := curDB.Limit(page.Size).Offset(page.GetStartPage()).Order("add_time desc").Find(&blogList) 
	return blogList, result.Error
}

func (blog *Blog) FindCountByTypeId() (count int64) {
	G.GLOBAL_DB.Model(blog).Where("id = ?", blog.TypeId).Count(&count)
	return
}

// Count the total page of the blog
func (blog *Blog) Count() (count int64) {
	G.GLOBAL_DB.Model(blog).Count(&count)
	return
}

func (blog *Blog) UpdataClick() (error) {
	db := G.GLOBAL_DB.Model(blog).
			Update("click_hit", gorm.Expr("click_hit + 1"))
	return db.Error
}

func (blog *Blog) UpdateReplay() (error) {
	db := G.GLOBAL_DB.Model(blog).
			Where("id = ? ", blog.Id).
			Update("replay_hit", gorm.Expr("replay_hit + ?", 1))
	return db.Error
}

// FindBlogComment : query the comment of the blog
func (blog *Blog) FindBlogComment() ([]Comment, error) {
	comments := make([]Comment, 0)
	db := G.GLOBAL_DB.Table("comment").
			Where("blog_id = ? and status = 0", blog.Id).
			Order("add_time desc").
			Find(&comments)
	if db.Error != nil {
		return nil, db.Error
	}
	return comments, nil
}

// BlogListWithTag
func (blog *Blog) BlogListWithTag(tag *Tag, page *utils.Page) ([]*Blog, error) {
	blogList := make([]*Blog, 0)
	db := G.GLOBAL_DB.Model(blog).
			Select("blog.id, title, typeid, add_time, update_time, click_hit, BLOG_TAG.TAG_ID AS TAG_ID").
			Joins("INNER JOIN BLOG_TAG ON BLOG.ID = BLOG_TAG.BLOG_ID").
			Where("tag_id = ", tag.Id).
			Limit(page.Size).
			Offset(page.GetStartPage()).
			Order("BLOG.ID ASC").
			Find(&blogList)
	
	if db.Error != nil {
		return nil, db.Error
	}
	return blogList, nil
}