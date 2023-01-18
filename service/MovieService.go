package service

import (
	"strconv"
	"strings"

	"github.com/rafaelmgr12/moviego-cli/model"
	"github.com/rafaelmgr12/moviego-cli/utils"
)

func GetAllMoviesFromCSV(filePath string) []model.Movie {
	records := utils.ReadCSVFile(filePath)

	var movies []model.Movie
	for _, record := range records[1:] {
		id, err := strconv.Atoi(record[0])
		if err != nil {
			panic(err)
		}
		movieNameSplits := strings.Split(record[1], "(")
		title := movieNameSplits[0]
		year := strings.Split(record[1], "(")[len(movieNameSplits)-1]

		yearInt, _ := strconv.Atoi(strings.Trim(year, ")"))

		movie := model.Movie{
			ID:    id,
			Title: title,
			Genre: record[2],
			Year:  yearInt,
		}

		movies = append(movies, movie)

	}
	return movies
}
