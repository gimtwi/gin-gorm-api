package repository

import (
	"errors"

	"github.com/gimtwi/gin-gorm-api/helper"
	"github.com/gimtwi/gin-gorm-api/model"
	"gorm.io/gorm"
)

type TagsRepositoryImpl struct {
	DB *gorm.DB
}

func NewTagsRepositoryImpl(db *gorm.DB) TagsRepository {
	return &TagsRepositoryImpl{DB: db}
}

func (t *TagsRepositoryImpl) Save(tags model.Tags) {
	result := t.DB.Create(&tags)
	helper.ErrorPanic(result.Error)
}

func (t *TagsRepositoryImpl) Update(tags model.Tags) {
	// insure if converting works or different types necessary at all
	var updateTag = model.UpdateTagsRequest(model.Tags{
		Id:   tags.Id,
		Name: tags.Name,
	})
	result := t.DB.Model(&tags).Updates(updateTag)
	helper.ErrorPanic(result.Error)
}

func (t *TagsRepositoryImpl) Delete(tagsId int) {
	var tag model.Tags
	result := t.DB.Where("id = ?", tagsId).Delete(&tag)
	helper.ErrorPanic(result.Error)
}

func (t *TagsRepositoryImpl) FindById(tagsId int) (tags model.Tags, err error) {
	var tag model.Tags
	result := t.DB.Find(&tag, tagsId)
	if result != nil {
		return tag, nil
	} else {
		return tags, errors.New("tag is not found")
	}
}

func (t *TagsRepositoryImpl) FindAll() []model.Tags {
	var tags []model.Tags
	result := t.DB.Find(&tags)
	helper.ErrorPanic(result.Error)
	return tags
}
