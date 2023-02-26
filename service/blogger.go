package service

import (
	"blog/server/initialize/global"
	"blog/server/model"
)

type BloggerService struct{}

var bloggerService BloggerService

func (b *BloggerService) Create(blogger *model.Blogger) error {
	return global.GLOBAL_DB.Create(blogger).Error
}

// Query the blogger with id 1
func (b *BloggerService) FindIdFirst() (*model.Blogger, error) {
	result := &model.Blogger{}
	err := global.GLOBAL_DB.Where("id = 1").First(result).Error
	return result, err
}

// Query the blogger by given name
func (b *BloggerService) FindByName(blogger *model.Blogger) (*model.Blogger, error) {
	result := &model.Blogger{}
	err := global.GLOBAL_DB.Where("username = ?", blogger.Username).First(result).Error
	return result, err
}

// Query for updating blogger password
func (b *BloggerService) UpdateSecurityInfo(blogger *model.Blogger) (*model.Blogger, error) {
	if blogger.Password == "" {
		err := global.GLOBAL_DB.Save(blogger).Error
		return blogger, err
	}
	err := global.GLOBAL_DB.Model(blogger).Update("password", blogger.Password).Error
	return blogger, err
}

// Query for update blogger information
func (b *BloggerService) UpdateInfo(blogger *model.Blogger) (*model.Blogger, error) {
	if blogger.Password == "" {
		err := global.GLOBAL_DB.Save(blogger).Error
		return blogger, err
	}
	err := global.GLOBAL_DB.Model(blogger).Updates(blogger).Error
	return blogger, err
}
