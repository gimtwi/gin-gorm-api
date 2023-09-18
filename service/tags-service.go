package service

import "github.com/gimtwi/gin-gorm-api/model"

type TagsService interface {
	Create(tags model.CreateTagsRequest)
	Update(tags model.UpdateTagsRequest)
	Delete(tagsId int)
	FindById(tagsId int) model.TagsResponse
	FindAll() []model.TagsResponse
}
