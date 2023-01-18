package database

import (
	"log"

	"github.com/rafaelmgr12/moviego-cli/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Connect() {
	connectionString := "host=localhost user=rafael password=secret dbname=moviesgo port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	DB, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Println(err)
	}
	DB.AutoMigrate(&model.Movie{})
}
