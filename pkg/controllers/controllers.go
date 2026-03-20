package controllers

import (
	"git.example.com/EnterpriseGradeSystem/pkg/models"
	"git.example.com/EnterpriseGradeSystem/pkg/services"
	"encoding/json"
	"log"
	"net/http"
)

func StartServer() {
	http.HandleFunc("/users", usersHandler)
	http.ListenAndServe(":8080", nil)
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	users := services.GetUsers()
	json.NewEncoder(w).Encode(users)
}