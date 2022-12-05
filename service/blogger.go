package service

import (
	G "blog/server/global"
	"blog/server/model"
)

type Blogger model.Blogger

func (b *Blogger) Find() (blogger *Blogger){
	blogger = new(Blogger)
	G.GLOBAL_DB.Where("id = 1").First(b)
	return
}