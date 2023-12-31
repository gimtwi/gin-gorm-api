package model

type Tags struct {
	Id   int    `gorm:"type:int;primary_key"`
	Name string `gorm:"type:varchar(255)"`
}

type CreateTagsRequest struct {
	Name string `validate:"required,min=1,max=200" json:"name"`
}

type UpdateTagsRequest struct {
	Id   int    `validate:"required"`
	Name string `validate:"required,min=1,max=200" json:"name"`
}

type TagsResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Response struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data,omitempty"`
}
