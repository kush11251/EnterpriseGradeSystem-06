package main

import (
	"git.example.com/EnterpriseGradeSystem/pkg/config"
	"git.example.com/EnterpriseGradeSystem/pkg/controllers"
	"log"
)

func main() {
	config.Init()
	controllers.StartServer()
	log.Println("Server started")
}