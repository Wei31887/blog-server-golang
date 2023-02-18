package model

import "time"

type Comment struct {
	Id        int    `gorm:"id" json:"id"`
	NickName  string `gorm:"nick_name" json:"nick_name"`
	Ip        string `gorm:"ip" json:"ip"`
	Content   string `gorm:"content" json:"content"`
	BlogId    int    `gorm:"blog_id" json:"blog_id"`
	Status    int    `gorm:"status" json:"status"`
	AddTime   time.Time `gorm:"add_time" json:"add_time"`
	BlogTitle string `gorm:"-" json:"blog_title"`
}

// Gorm 約定 table name
func (Comment) TableName() string {
	return "comment"
}