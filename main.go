package main

import (
	"log"
	"os"
	"sync"

	"github.com/rafaelmgr12/moviego-cli/controller"
	"github.com/rafaelmgr12/moviego-cli/database"
	"github.com/rafaelmgr12/moviego-cli/service"
	"github.com/urfave/cli"
)

var app = cli.NewApp()
var wg sync.WaitGroup

func info() {
	app.Name = "Moviego CLI"
	app.Usage = "A CLI for inserting data into a database using a CSV file"
	app.Author = "rafaelmgr12"
	app.Version = "1.0.0"
}
func commands() {
	app.Commands = []cli.Command{
		{
			Name:    "insert",
			Aliases: []string{"i"},
			Usage:   "Add a movie to the database",
			Action: func(c *cli.Context) {
				database.Connect()
				movies := service.GetAllMoviesFromCSV("./input/movies.csv")
				for _, movie := range movies {
					go controller.SaveMovieInDB(movie)
					wg.Add(1)
				}
				wg.Wait()

			},
		},
	}
}

func main() {
	info()
	commands()

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
