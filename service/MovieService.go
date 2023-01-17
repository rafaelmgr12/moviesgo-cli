package service

import (
	"github.com/rafaelmgr12/moviego-cli/model"
	"github.com/rafaelmgr12/moviego-cli/utils"
)

func GetAllMoviesFromCSV(filePath string) []model.Movie {
	records := utils.ReadCSVFile(filePath)
	var movies []model.Movie
	for _, record := range records {
		movie := model.Movie{
			ID:    record[0],
			Title: record[1],
			Genre: record[2],
		}
		movies = append(movies, movie)
	}
	return movies
}
