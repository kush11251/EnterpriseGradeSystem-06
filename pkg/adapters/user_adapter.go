package adapters

import (
	"git.example.com/EnterpriseGradeSystem/pkg/models"
)

type UserAdapter struct {}

func (u *UserAdapter) GetUser(id string) models.User {
	// Implement adapter logic here
	return models.User{ID: id, Name: "John"}
}