package service

import (
	G "blog/server/global"
	"blog/server/model"
)

type Blogger model.Blogger

// Gorm 約定 table name
func (Blogger) TableName() string {
	return "blogger"
}

// Query the blogger with id 1
func (b *Blogger) FindIdFirst() (*Blogger, error){
	blogger := new(Blogger)
	db := G.GLOBAL_DB.Where("id = 1").First(blogger)
	if db.Error != nil {
		return nil, db.Error
	}
	return blogger, nil
}

// Query the blogger by given name
func (b *Blogger) FindByName() (*Blogger, error) {
	blogger := new(Blogger)
	db := G.GLOBAL_DB.Where("username = ?", b.Username).First(blogger)
	if db.Error != nil {
		return nil, db.Error
	}
	return blogger, nil
}