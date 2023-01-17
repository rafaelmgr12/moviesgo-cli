package main

import (
	"testing"

	"github.com/rafaelmgr12/moviego-cli/database"
	"github.com/rafaelmgr12/moviego-cli/service"
	"github.com/stretchr/testify/assert"
)

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

func TestSaveMovieInDB(t *testing.T) {

}
