<h1 align="center">
  <img alt="Logo" src="./img/logo.jfif">
</h1>

<h1 align="center">MoviesGo</h1>
<p align = "center">A CLI to insert data into Postgres from a csv file</p>

[]()
<p align="center">
  <a href="#-technology">Technology</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
    <a href="#-project">Project</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
  <a href="#-how-to-run">How to Run</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
  <a href="#-license">License</a>
</p>

<p align="center">
  <img alt="License" src="https://img.shields.io/static/v1?label=license&message=MIT&color=8257E5&labelColor=000000">
</p>

## Introduction
This project is based on the following challenge, where we have a CSV file, and we have to insert the data into a database. The challenge has just a few requirements, but it can be implemented in any language.

The following challenge can be found [here](https://app.devgym.com.br/challenges/ec36e7e2-6a2d-4406-98e1-3029f843b5c3).

## ✨ Technology

The Project was develop as using the following techs:
- [Go](https://go.dev/)
- [Gorm](https://github.com/go-gorm/gorm)
- [Postgres](https://www.postgresql.org/)
- [cli](https://github.com/urfave/cli)



## 💻 Project
In this challenge, you must create a command line program (CLI) that reads a CSV file of movies and populates a database, thinking about performance and hoping that the file can grow considerably


###  📓 Requirements 
Each line of the csv contains columns that must be saved in separate columns/fields in the database:
* ID - integer that identifies the movie found in the 1st column of the csv
* title - movie title found in the 2nd column of the CSS. With the value "Jumanji (1995)" the title is Jumanji
* "Year - Year of the movie found in the 2nd column of the CSV. With the value "Jumanji (1995)" the year is 1995
* Genres - multiple values with movie genres separated by |. Found in the 3rd column.

The script must think about performance and take advantage of concurrency/parallelism to populate the database.

Download the csv file [here](https://github.com/devgymbr/files/raw/main/movie.csv.zip).


## 🚀 How to Run

To run this project, you first need to set up a database, PostgreSQL and look in the `database.go` file to see what settings are needed. 

Having done this, we have:
1. go build main.go
2. ./main.go i

## 📄 License
The projects is under the MIT license. See the file [LICENSE](LICENSE) fore more details

---
## Author

Made with ♥ by Rafael 👋🏻


[![Linkedin Badge](https://img.shields.io/badge/-Rafael-blue?style=flat-square&logo=Linkedin&logoColor=white&link=https://www.linkedin.com/in/tgmarinho/)](https://www.linkedin.com/in/rafael-mgr/)
[![Gmail Badge](https://img.shields.io/badge/-Gmail-red?style=flat-square&link=mailto:nelsonsantosaraujo@hotmail.com)](mailto:ribeirorafaelmatehus@gmail.com)