package service

import (
	G "blog/server/global"
	"blog/server/model"
	"blog/server/utils"
)

type BlogType model.BlogType

// Gorm 約定 table name
func (BlogType) TableName() string {
	return "blog_type"
}

func (b *BlogType) Create() error {
	db := G.GLOBAL_DB.Create(b)
	return db.Error
}

func (b *BlogType) Update() error {
	db := G.GLOBAL_DB.Save(b)
	return db.Error
}

func (b *BlogType) Delete() error {
	db := G.GLOBAL_DB.Delete(b)
	return db.Error	
}

func (b *BlogType) FindTypeAll() ([]*BlogType, error) {
	var blogTypes = make([]*BlogType, 0)
	if db := G.GLOBAL_DB.Order("sort asc").Find(&blogTypes); db.Error != nil {
		return nil, db.Error
	}
	return blogTypes, nil
}

func (b *BlogType) FindTypeIdOne() (*BlogType, error){
	var blogType = new(BlogType)
	db := G.GLOBAL_DB.Where("id = ?", b.Id).First(&blogType)
	if db.Error != nil {
		return nil, db.Error
	}
	return blogType, nil
}

// FindTypeCount : query the count of different types
func (b *BlogType) FindAllTypeCount() ([]map[string]interface{}, error) {
	typeMaps := make([]map[string]interface{}, 0)
	sql := `select blog_type.id, count(blog.typeid) as b_count, blog_type.name as b_name from blog 
			left join blog_type on blog.typeid = blog_type.id 
			group by blog_type.id order by id; `

	result, err := G.GLOBAL_DB.Raw(sql).Rows()
	if err != nil {
		return nil, err
	}
	defer result.Close()
	
	for result.Next() {
		var (
			id int
			bCount int
			bName string
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

func (b *BlogType) FindTypeCount() (int, error) {
	var count int64
	if db := G.GLOBAL_DB.Model(b).Count(&count); db.Error != nil {
		return 0, db.Error
	}
	return int(count), nil
}

func (b *BlogType) FindTypeList(page utils.Page) ([]*BlogType, error) {
	blogTypes := make([]*BlogType, 0)
	db := G.GLOBAL_DB.Model(b).Limit(page.Size).Offset(page.GetStartPage()).Order("sort asc").Find(&blogTypes)
	if db.Error != nil {
		return nil, db.Error
	}
	return blogTypes, nil
} 
