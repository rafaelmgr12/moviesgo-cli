package service

import (
	"strconv"

	"github.com/rafaelmgr12/moviego-cli/model"
	"github.com/rafaelmgr12/moviego-cli/utils"
)

func GetAllMoviesFromCSV(filePath string) []model.Movie {
	records := utils.ReadCSVFile(filePath)
	nameExistMap := make(map[string]bool)

	var movies []model.Movie
	for _, record := range records[1:] {
		id, err := strconv.Atoi(record[0])
		if err != nil {
			panic(err)
		}
		movie := model.Movie{
			ID:    id,
			Title: record[1],
			Genre: record[2],
		}
		if _, ok := nameExistMap[movie.Title]; ok {
			continue
		} else {
			movies = append(movies, movie)
		}
	}
	return movies
}
