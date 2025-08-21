package user

import (
	"fmt"
	"strconv"
	"task_csv/services/csv"
)

type void struct{}

func ParseBlacklistCSV(filename string) (map[int]void, error) {
	blacklist, err := csv.Read(filename)

	if err != nil {
		return nil, fmt.Errorf("error parsing blacklist: %w", err)
	}

	results := make(map[int]void, len(blacklist))

	for _, v := range blacklist[1:] {
		id, err := strconv.Atoi(v[0])

		if err != nil {
			continue
		}

		results[id] = void{}
	}

	return results, nil
}
