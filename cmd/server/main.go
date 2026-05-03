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
	// เชื่อมต่อกับ database
	config.ConnectDatabase()

	// Setup Slot module
	slotRepo := &repository.SlotRepository{}
	slotService := &service.SlotService{Repo: slotRepo}
	handler.SetupSlotHandler(slotService)

	// Setup Doctor module
	doctorRepo := &repository.DoctorRepository{}
	doctorService := &service.DoctorService{Repo: doctorRepo}
	handler.SetupDoctorHandler(doctorService)

	// Gin
	r := gin.Default()

	// Register routes
	route.RegisterSlotRoutes(r)
	route.RegisterDoctorRoutes(r)

	r.Run(":8080")
}