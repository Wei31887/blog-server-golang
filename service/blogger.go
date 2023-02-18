package service

import (
	G "blog/server/global"
	"blog/server/model"
)

type BloggerService struct {}

func (b *BloggerService) Create(blogger *model.Blogger) (*model.Blogger, error) {
	err := G.GLOBAL_DB.Debug().Create(blogger).Error
	return blogger, err
}


// Query the blogger with id 1
func (b *BloggerService) FindIdFirst() (*model.Blogger, error){
	result := &model.Blogger{}
	err := G.GLOBAL_DB.Where("id = 1").First(result).Error
	return result, err
}

// Query the blogger by given name
func (b *BloggerService) FindByName(blogger *model.Blogger) (*model.Blogger, error) {
	result := &model.Blogger{}
	err := G.GLOBAL_DB.Where("username = ?", blogger.Username).First(result).Error
	return result, err
}

// Query for updating blogger password
func (b *BloggerService) UpdateSecurityInfo(blogger *model.Blogger) (*model.Blogger, error) {
	if blogger.Password == "" {
		err := G.GLOBAL_DB.Save(b).Error
		return blogger, err
	} 
	err := G.GLOBAL_DB.Model(b).Update("password", blogger.Password).Error
	return blogger, err
}

// Query for update blogger information
func (b *BloggerService) UpdateInfo(blogger *model.Blogger) (*model.Blogger, error) {
	if blogger.Password == "" {
		err := G.GLOBAL_DB.Save(b).Error
		return blogger, err
	} 
	err := G.GLOBAL_DB.Model(b).Updates(blogger).Error
	return blogger, err
}