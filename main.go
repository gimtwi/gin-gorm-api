package main

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"

	"github.com/gimtwi/gin-gorm-api/config"
	"github.com/gimtwi/gin-gorm-api/controller"
	"github.com/gimtwi/gin-gorm-api/helper"
	"github.com/gimtwi/gin-gorm-api/model"
	"github.com/gimtwi/gin-gorm-api/repository"
	"github.com/gimtwi/gin-gorm-api/router"
	"github.com/gimtwi/gin-gorm-api/service"
)

func main() {
	log.Info().Msg("Server is running!")

	db := config.DatabaseConnection()
	validate := validator.New()

	db.Table("tags").AutoMigrate(&model.Tags{})

	tagsRepo := repository.NewTagsRepositoryImpl(db)
	tagsService := service.NewTagsServiceImpl(tagsRepo, validate)
	tagsController := controller.NewTagsController(tagsService)

	routes := router.NewRouter(tagsController)

	server := &http.Server{
		Addr:    ":8888",
		Handler: routes,
	}

	err := server.ListenAndServe()
	helper.ErrorPanic(err)
}
