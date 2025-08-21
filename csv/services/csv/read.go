package csv

import (
	"encoding/csv"
	"log"
	"os"
)

func Read(filename string) ([][]string) {
	f, err := os.Open(filename)

	if err != nil {
		log.Fatal("Unable to read file " + filename, err)
	}
	defer f.Close()

	reader := csv.NewReader(f)
	records, err := reader.ReadAll()

	if err != nil {
		log.Fatal("Unable to parse file as CSV for " + filename, err)
	}

	return records
}