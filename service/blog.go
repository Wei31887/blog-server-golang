package service

import (
	. "blog/server/global"
	"blog/server/model"
	"blog/server/utils"

	"gorm.io/gorm"
)

type Blog model.Blog

// FindNextBlog : query the next blog of given blog
func (blog *Blog) FindNextBlog() (*Blog, error) {
	resBlog := new(Blog)
	db := GLOBAL_DB.Where("id > ?", blog.Id).First(resBlog)
	if db.Error != nil {
		return nil, db.Error
	}
	return resBlog, nil
}

// FindPreviosBlog : query the previous blog of given blog
func (blog *Blog) FindPreviousBlog() (*Blog, error) {
	resBlog := new(Blog)
	db := GLOBAL_DB.Where("id < ?", blog.Id).Order("id desc").First(resBlog)
	if db.Error != nil {
		return nil, db.Error
	}
	return resBlog, nil
}

// FindBlogWithTypeName : query the blog with type name by the given blog id
func (blog *Blog) FindBlogWithTypeName() (*Blog, error) {
	resBlog := new(Blog)
	db := GLOBAL_DB.Table("blog").Select("*, blog_type.name as type_name").
					Joins("left join blog_type on blog.type_id = blog_type.id").
					Where("id = ?", blog.Id).Order("blog_type.sort asc").Find(resBlog)
	if db.Error != nil {
		return nil, db.Error
	}
	return resBlog, nil
}


// FindList : query the blog list of the page
func (blog *Blog) FindList(page utils.Page) ([]*Blog, error) {
	blogList := make([]*Blog, 0)
	curDB := GLOBAL_DB.Table("blog").Select("*, blog_type.name as type_name").
				Joins("left join blog_type on blog.typeid = blog_type.id")
	
	if blog.TypeId > 0 {
		curDB = curDB.Where("blog.typeid = ? ", blog.TypeId)
	}
	
	// Limit the maximum query number and offset
	result := curDB.Limit(page.Size).Offset(page.GetStartPage()).Order("add_time asc").Find(&blogList) 
	return blogList, result.Error
}

func (blog *Blog) FindCountByTypeId() (count int64) {
	GLOBAL_DB.Model(blog).Where("id = ?", blog.TypeId).Count(&count)
	return
}

func (blog *Blog) Count() (count int64) {
	GLOBAL_DB.Model(blog).Count(&count)
	return
}

func (blog *Blog) UpdataClick() (error) {
	db := GLOBAL_DB.Model(blog).Where("id = ?", blog.Id).Update("click_hit", gorm.Expr("click + 1"))
	return db.Error
}

// FindBlogComment : query the comment of the blog
func (blog *Blog) FindBlogComment() ([]Comment, error) {
	comments := make([]Comment, 0)
	db := GLOBAL_DB.Table("comment").Where("blog_id = ? and status = 1", blog.Id).
			Order("add_time asc").Find(&comments)
	if db.Error != nil {
		return nil, db.Error
	}
	return comments, nil
}