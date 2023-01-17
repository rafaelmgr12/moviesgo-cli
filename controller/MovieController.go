package controller

import (
	"github.com/rafaelmgr12/moviego-cli/database"
	"github.com/rafaelmgr12/moviego-cli/model"
)

func SaveMovieInDB(movie model.Movie) {
	database.Connect()
	database.DB.Create(&movie)

}
