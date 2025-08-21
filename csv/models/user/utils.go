package user

import (
	"fmt"
	"strconv"
	"strings"
	"task_csv/services/csv"
	"task_csv/shared/validator"
)

func ParseUsersFromCSV(filename string) ([]UserCSV, error) {
	records, err := csv.Read(filename)

	if err != nil {
		return nil, fmt.Errorf("error parsing users: %w", err)
	}

	results := make([]UserCSV, 0, len(records))

	for _, value := range records[1:] {
		id, _ := strconv.Atoi(value[0])

		results = append(results, UserCSV{
			ID:     id,
			Name:   strings.ToLower(value[1]),
			Email:  value[2],
			RoleID: value[3],
		})
	}

	return results, nil
}

func ParseRoleMapFromCSV(rolesFileName string) (map[string]string, error) {
	roles, err := csv.Read(rolesFileName)

	if err != nil {
		return nil, fmt.Errorf("error parsing roles: %w", err)
	}

	roleMap := make(map[string]string)

	for _, value := range roles[1:] {
		roleMap[value[0]] = value[1]
	}

	return roleMap, nil
}

func UserCSVToUserWithRoles(users []UserCSV, roleMap map[string]string) []User {
	results := make([]User, 0, len(users))

	for i := range users {
		if roleName, exists := roleMap[users[i].RoleID]; exists {
			results = append(results, User{
				ID:       users[i].ID,
				Name:     users[i].Name,
				Email:    users[i].Email,
				RoleName: roleName,
			})
		}

	}

	return results
}

func GetUsersByRole(users []User, roleName string) []User {
	results := make([]User, 0)

	for _, user := range users {
		if user.RoleName == roleName {
			results = append(results, user)
		}
	}

	return results
}

func ValidateUsers(users []User, blacklist map[int]void) []User {
	results := make([]User, 0)

	for _, user := range users {
		_, isBlacklisted := blacklist[user.ID]
		if validator.IsValidEmail(user.Email) && !isBlacklisted {
			results = append(results, user)
		}
	}

	return results
}
