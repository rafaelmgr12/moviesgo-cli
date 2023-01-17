package database

import (
	"log"

	"github.com/rafaelmgr12/moviego-cli/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Connect() {
	connectionString := "root:root@tcp(127.0.0.1:3306)/moviesgo?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Println(err)
	}
	DB.AutoMigrate(&model.Movie{})
}
