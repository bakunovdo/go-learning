package user

import (
	"strconv"
	"task_csv/services/csv"
)

func ParseBlacklistCSV(filename string) map[int]bool {
	blacklist := csv.Read(filename)

	results := make(map[int]bool, 0)

	for _, v := range blacklist[1:] {
		id, _ := strconv.Atoi(v[0])
		results[id] = true
	}

	return results
}