package user

import (
	"fmt"
	"strconv"
	"task_csv/services/csv"
)

func ParseUsersFromCSV(filename string) ([]User) {
	records := csv.Read(filename)
	results := make([]User, 0, len(records))

	for _, value := range records[1:] {
		id, _ := strconv.Atoi(value[0])

		results = append(results, User{
				ID: id,
				Name: value[1],
				Email: value[2],
				RoleName: value[3],
		})
	}

	return results
}

func EnrichWithUserRoles(users []User, rolesFileName string) []User {
	roles := csv.Read(rolesFileName)
	roleMap := make(map[string]string)

	for _, value := range roles[1:] {
		roleMap[value[0]] = value[1]
	}

	fmt.Println(roleMap)

	for i := range users {
		if roleName, exists := roleMap[users[i].RoleName]; exists {
			users[i].RoleName = roleName
		}
	}

	return users
}