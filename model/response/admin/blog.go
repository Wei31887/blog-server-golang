package adminresponse

import "blog/server/model"


type BlogFindOneResponse struct {
	Blog *model.Blog  `json:"blog"`
	Tags []*model.Tag `json:"tags"`
}