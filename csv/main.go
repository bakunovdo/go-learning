package main

import (
	"fmt"
	"task_csv/models/user"
)

func handleError(err error, message string) {
	if err != nil {
		fmt.Printf("error %s: %v\n", message, err)
	}
}

func main() {
	usersCSV, err := user.ParseUsersFromCSV("./assets/users.csv")
	handleError(err, "error parsing users")

	rolesMap, err := user.ParseRoleMapFromCSV("./assets/user_roles.csv")
	handleError(err, "error parsing roles")

	blacklist, err := user.ParseBlacklistCSV("./assets/blacklist.csv")
	handleError(err, "error parsing blacklist")

	users := user.UserCSVToUserWithRoles(usersCSV, rolesMap)
	validatedUsers := user.ValidateUsers(users, blacklist)

	admins := user.GetUsersByRole(validatedUsers, "admin")
	members := user.GetUsersByRole(validatedUsers, "member")
	guests := user.GetUsersByRole(validatedUsers, "guest")

	fmt.Printf("admins: %v\n", admins)
	fmt.Printf("members: %v\n", members)
	fmt.Printf("guests: %v\n", guests)
}
