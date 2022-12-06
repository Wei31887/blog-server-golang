package service

import (
	G "blog/server/global"
	"blog/server/model"
)

type BlogType model.BlogType

// Gorm 約定 table name
func (BlogType) TableName() string {
	return "blog_type"
}

// FindTypeCount : query the count of different types
func (b *BlogType) FindTypeCount() ([]map[string]interface{}, error) {
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