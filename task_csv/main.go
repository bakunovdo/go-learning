package main

import (
	"fmt"
	"task_csv/models/user"
	"task_csv/shared/validator"
)


func main() {
	results := make([]user.User, 0)

	users := user.EnrichWithUserRoles(user.ParseUsersFromCSV("./assets/users.csv"), "./assets/user_roles.csv")
	blacklist := user.ParseBlacklistCSV("./assets/blacklist.csv")

	for _, user := range(users) {
		_, isBlacklisted := blacklist[user.ID]
		if validator.IsValidEmail(user.Email) && !isBlacklisted {
			results = append(results, user)
		}
	}

	fmt.Println(results)
}