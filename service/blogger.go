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

func (b *Blogger) Create() (*Blogger, error) {
	blogger := new(Blogger)
	db := G.GLOBAL_DB.Create(blogger)
	if db.Error != nil {
		return nil, db.Error
	}
	return blogger, nil
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

// Query for updating blogger password
func (b *Blogger) UpdatePassword() (error) {
	if b.Password == "" {
		db := G.GLOBAL_DB.Save(b)
		if db.Error != nil {
			return db.Error
		}
	} 
	db := G.GLOBAL_DB.Model(b).Update("password", b.Password)
	if db.Error != nil {
		return db.Error
	}
	return nil
}

// Query for update blogger information
func (b *Blogger) UpdateInfo() (error) {
	if b.Password == "" {
		db := G.GLOBAL_DB.Save(b)
		if db.Error != nil {
			return db.Error
		}
	} 
	db := G.GLOBAL_DB.Model(b).Updates(b)
	if db.Error != nil {
		return db.Error
	}
	return nil
}