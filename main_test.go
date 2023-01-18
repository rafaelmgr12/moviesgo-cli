package main

import (
	"testing"

	"github.com/rafaelmgr12/moviego-cli/database"
	"github.com/rafaelmgr12/moviego-cli/model"
	"github.com/rafaelmgr12/moviego-cli/service"
	"github.com/stretchr/testify/assert"
)

var IDTest int

func MockMovie() {
	database.Connect()
	movie := model.Movie{ID: 1, Title: "Test", Genre: "Test", Year: 2021}
	database.DB.Create(&movie)
	IDTest = movie.ID

}

func DeleteMovieMock() {
	var movie model.Movie
	database.Connect()
	database.DB.First(&movie)
	database.DB.Delete(&movie, IDTest)
}

func TestExample(t *testing.T) {
	assert.Equal(t, 1, 1, "Should be equal")
}

func TestGetAllMoviesFromCSV(t *testing.T) {
	movies := service.GetAllMoviesFromCSV("./input/movies.csv")
	assert.Equal(t, 27279, len(movies), "Should be equal")

}

func TestDatabaseConnection(t *testing.T) {
	database.Connect()
	assert.NotNil(t, database.DB, "Should be not nil")
}

func TestDatabaseInsert(t *testing.T) {
	MockMovie()
	defer DeleteMovieMock()
	var movie model.Movie
	result := database.DB.First(&movie, IDTest)
	assert.Equal(t, int(result.RowsAffected), 1, "Should return the one row with the test value")
}
