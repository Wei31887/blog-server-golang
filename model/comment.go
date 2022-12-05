package model

type Comment struct {
	Id        int    `gorm:"id" json:"id"`
	Ip        string `gorm:"ip" json:"ip"`
	Content   string `gorm:"content" json:"content"`
	BlogId    int    `gorm:"blog_id" json:"blog_id"`
	Status    int    `gorm:"status" json:"status"`
	AddTime   string `gorm:"add_time" json:"add_time"`
	BlogTitle string `gorm:"-" json:"blog_title"`
}