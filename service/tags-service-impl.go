package service

import (
	"github.com/gimtwi/gin-gorm-api/helper"
	"github.com/gimtwi/gin-gorm-api/model"
	"github.com/gimtwi/gin-gorm-api/repository"
	"github.com/go-playground/validator/v10"
)

type TagsServiceImpl struct {
	TagsRepository repository.TagsRepository
	Validate       *validator.Validate
}

func NewTagsServiceImpl(tagRepo repository.TagsRepository, validate *validator.Validate) TagsService {
	return &TagsServiceImpl{
		TagsRepository: tagRepo,
		Validate:       validate,
	}
}

func (t *TagsServiceImpl) Create(tags model.CreateTagsRequest) {
	err := t.Validate.Struct(tags)
	helper.ErrorPanic(err)
	tagModel := model.Tags{
		Name: tags.Name,
	}
	t.TagsRepository.Save(tagModel)
}

func (t *TagsServiceImpl) Delete(tagsId int) {
	t.TagsRepository.Delete(tagsId)
}

func (t *TagsServiceImpl) FindAll() []model.TagsResponse {
	result := t.TagsRepository.FindAll()

	var tags []model.TagsResponse
	for _, value := range result {
		tag := model.TagsResponse(model.Tags{
			Id:   value.Id,
			Name: value.Name,
		})
		tags = append(tags, tag)
	}
	return tags
}

func (t *TagsServiceImpl) FindById(tagsId int) model.TagsResponse {
	tagData, err := t.TagsRepository.FindById(tagsId)
	helper.ErrorPanic(err)

	tagRes := model.TagsResponse(model.Tags{
		Id:   tagData.Id,
		Name: tagData.Name,
	})
	return tagRes
}

func (t *TagsServiceImpl) Update(tags model.UpdateTagsRequest) {
	tagData, err := t.TagsRepository.FindById(tags.Id)
	helper.ErrorPanic(err)

	tagData.Name = tags.Name
	t.TagsRepository.Update(tagData)
}
