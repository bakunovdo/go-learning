package csv

import (
	"encoding/csv"
	"fmt"
	"os"
)

func Read(filename string) ([][]string, error) {
	f, err := os.Open(filename)

	if err != nil {
		return nil, fmt.Errorf("open file: %w", err)
	}

	defer f.Close()

	reader := csv.NewReader(f)
	records, err := reader.ReadAll()

	if err != nil {
		return nil, fmt.Errorf("unable to parse file as CSV for %s: %w", filename, err)
	}

	return records, nil
}
