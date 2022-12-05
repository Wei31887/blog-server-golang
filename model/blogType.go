package model

type BlogType struct {
	Id   int    `gorm:"id" json:"id"`
	Name string `gorm:"name" json:"name"`
	Sort int    `gorm:"sort" json:"sort"`
}