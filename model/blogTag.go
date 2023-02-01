package model

type Tag struct {
	Id		int	 	`gorm:"id" json:"id"`
	TagName string 	`gorm:"tag_name" json:"tag_name"`
	Sort 	int    	`gorm:"sort" json:"sort"`
	Count 	int 	`gorm:"->" json:"count"`	
}

type BlogTag struct {
	Id 		int		`gorm:"id" json:"id"`
	BlogId 	int 	`gorm:"blog_name" json:"blog_name"`
	TagId 	int 	`gorm:"tag_id" json:"tag_id"`
}