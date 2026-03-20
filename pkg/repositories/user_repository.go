package repositories

import (
	"git.example.com/EnterpriseGradeSystem/pkg/models"
)

func GetUsers() []models.User {
	// Implement database logic here
	return []models.User{{ID: "1", Name: "John"}, {ID: "2", Name: "Jane"}}
}