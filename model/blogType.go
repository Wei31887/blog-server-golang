package model

type BlogType struct {
	Id   int    `gorm:"id" json:"id"`
	Name string `gorm:"name" json:"name"`
	Sort int    `gorm:"sort" json:"sort"`
}

// Gorm 約定 table name
func (BlogType) TableName() string {
	return "blog_type"
}