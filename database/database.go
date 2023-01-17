package database

import (
	"github.com/rafaelmgr12/moviego-cli/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Connect() {
	connectionString := "host=localhost user=root password=root dbname=moviesgo port=3306 sslmode=disable"
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}
	DB.AutoMigrate(&model.Movie{})
}
