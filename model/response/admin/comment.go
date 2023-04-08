package adminresponse

import "time"

type CommentListResponse struct {
	Id       int       `gorm:"id" json:"id"`
	NickName string    `gorm:"nick_name" json:"nick_name"`
	Ip       string    `gorm:"ip" json:"ip"`
	Content  string    `gorm:"content" json:"content"`
	BlogId   int       `gorm:"blog_id" json:"blog_id"`
	Status   int       `gorm:"status" json:"status"`
	AddTime  time.Time `gorm:"add_time" json:"add_time"`
	Title    string    `gorm:"title" json:"blog_title"`
}
