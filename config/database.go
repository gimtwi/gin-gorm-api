package config

import (
	"fmt"

	"github.com/gimtwi/gin-gorm-api/helper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// TODO make it an environmental variable
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbName   = "postgres"
)

func DatabaseConnection() *gorm.DB {
	sqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
	db, err := gorm.Open(postgres.Open(sqlInfo), &gorm.Config{})
	helper.ErrorPanic(err)

	return db
}
