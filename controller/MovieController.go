package controller

import (
	"sync"

	"github.com/rafaelmgr12/moviego-cli/database"
	"github.com/rafaelmgr12/moviego-cli/model"
)

var mut sync.Mutex
var wg sync.WaitGroup

func SaveMovieInDB(movie model.Movie) {

	mut.Lock()
	defer mut.Unlock()
	database.DB.Create(&movie)

}
