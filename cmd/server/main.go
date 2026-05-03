package main

import (
	"telemedicine-api/config"
	"telemedicine-api/handler"
	"telemedicine-api/repository"
	"telemedicine-api/route"
	"telemedicine-api/service"

	"github.com/gin-gonic/gin"
)

func main() {

	config.ConnectDatabase()

	// Setup Doctor module
	doctorRepo := &repository.DoctorRepository{}
	doctorService := &service.DoctorService{Repo: doctorRepo}
	handler.SetupDoctorHandler(doctorService)

	r := gin.Default()

	route.RegisterDoctorRoutes(r)

	r.Run(":8080")
}
