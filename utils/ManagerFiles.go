package utils

import (
	"bufio"
	"encoding/csv"
	"log"
	"os"
)

func ReadCSVFile(filename string) [][]string {
	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	reader := csv.NewReader(bufio.NewReader(file))
	reader.LazyQuotes = true
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	return records
}
