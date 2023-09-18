package controller

import (
	"net/http"
	"strconv"

	"github.com/gimtwi/gin-gorm-api/helper"
	"github.com/gimtwi/gin-gorm-api/model"
	"github.com/gimtwi/gin-gorm-api/service"
	"github.com/gin-gonic/gin"
)

type TagsController struct {
	TagsService service.TagsService
}

func NewTagsController(service service.TagsService) *TagsController {
	return &TagsController{
		TagsService: service,
	}
}

func (controller *TagsController) Create(ctx *gin.Context) {
	createTagReq := model.CreateTagsRequest{}
	err := ctx.ShouldBindJSON(&createTagReq)
	helper.ErrorPanic(err)

	controller.TagsService.Create(createTagReq)
	webRes := model.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webRes)
}

func (controller *TagsController) Update(ctx *gin.Context) {
	updateTagsReq := model.UpdateTagsRequest{}
	err := ctx.ShouldBindJSON(&updateTagsReq)
	helper.ErrorPanic(err)

	tagId := ctx.Param("tagId")
	id, err := strconv.Atoi(tagId)
	helper.ErrorPanic(err)
	updateTagsReq.Id = id

	controller.TagsService.Update(updateTagsReq)
	webRes := model.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webRes)
}

func (controller *TagsController) Delete(ctx *gin.Context) {
	tagId := ctx.Param("tagId")
	id, err := strconv.Atoi(tagId)
	helper.ErrorPanic(err)

	controller.TagsService.Delete(id)
	webRes := model.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webRes)
}

func (controller *TagsController) FindById(ctx *gin.Context) {
	tagId := ctx.Param("tagId")
	id, err := strconv.Atoi(tagId)
	helper.ErrorPanic(err)

	tagRes := controller.TagsService.FindById(id)
	webRes := model.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   tagRes,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webRes)
}

func (controller *TagsController) FindAll(ctx *gin.Context) {
	tagsRes := controller.TagsService.FindAll()
	webRes := model.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   tagsRes,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webRes)
}
