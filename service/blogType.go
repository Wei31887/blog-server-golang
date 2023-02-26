package service

import (
	"blog/server/initialize/global"
	"blog/server/model"
	"blog/server/utils"
)

type BlogTypeService struct{}

func (b *BlogTypeService) Create(blogType *model.BlogType) error {
	return global.GLOBAL_DB.Create(blogType).Error
}

func (b *BlogTypeService) Update(blogType *model.BlogType) error {
	return global.GLOBAL_DB.Save(blogType).Error
}

func (b *BlogTypeService) Delete(blogType *model.BlogType) error {
	return global.GLOBAL_DB.Delete(blogType).Error
}

func (b *BlogTypeService) FindTypeAll() ([]*model.BlogType, error) {
	var blogTypes = make([]*model.BlogType, 0)
	err := global.GLOBAL_DB.Order("sort asc").Find(&blogTypes).Error
	return blogTypes, err
}

// FindTypeIdOne : returns the blog type by id
func (b *BlogTypeService) FindTypeIdOne(blogType *model.BlogType) (*model.BlogType, error) {
	var result = &model.BlogType{}
	err := global.GLOBAL_DB.Where("id = ?", blogType.Id).First(result).Error
	return result, err
}

// FindTypeCount : query the count of different types
func (b *BlogTypeService) FindAllTypeCount() ([]map[string]interface{}, error) {
	typeMaps := make([]map[string]interface{}, 0)
	sql := `select blog_type.id, count(blog.type_id) as b_count, blog_type.name as b_name from blog 
			left join blog_type on blog.type_id = blog_type.id 
			group by blog_type.id order by id; `

	result, err := global.GLOBAL_DB.Raw(sql).Rows()
	if err != nil {
		return nil, err
	}
	defer result.Close()

	for result.Next() {
		var (
			id     int
			bCount int
			bName  string
		)
		_ = result.Scan(&id, &bCount, &bName)

		typeMap := make(map[string]interface{})
		typeMap["type_id"] = id
		typeMap["b_count"] = bCount
		typeMap["b_name"] = bName
		typeMaps = append(typeMaps, typeMap)
	}
	return typeMaps, nil
}

func (b *BlogTypeService) FindTypeCount() (int, error) {
	var count int64
	err := global.GLOBAL_DB.Model(&model.BlogType{}).Count(&count).Error
	return int(count), err
}

func (b *BlogTypeService) FindTypeList(page utils.Page) ([]*model.BlogType, error) {
	blogTypes := make([]*model.BlogType, 0)
	err := global.GLOBAL_DB.Model(&model.BlogType{}).
		Limit(page.Size).
		Offset(page.GetStartPage()).
		Order("sort asc").
		Find(&blogTypes).Error
	return blogTypes, err
}
