package services

import (
	"git.example.com/EnterpriseGradeSystem/pkg/models"
	"git.example.com/EnterpriseGradeSystem/pkg/repositories"
)

func GetUsers() []models.User {
	return repositories.GetUsers()
}