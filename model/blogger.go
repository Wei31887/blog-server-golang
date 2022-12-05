package model

type Blogger struct {
	Id       int    `gorm:"id" json:"id"`
	Username string `gorm:"username" json:"username"` //
	Password string `gorm:"password" json:"password"` //
	Nickname string `gorm:"nickname" json:"nickname"` //
	Sign     string `gorm:"sign" json:"sign"`         //
	Profile  string `gorm:"profile" json:"profile"`   //
	Img      string `gorm:"img" json:"img"`           //
}