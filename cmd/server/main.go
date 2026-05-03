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
	// เชื่อต่อกับ database
	config.ConnectDatabase()

	// Setup Slot module
	slotRepo := &repository.SlotRepository{}
	slotService := &service.SlotService{Repo: slotRepo}
	handler.SetupSlotHandler(slotService)

	// Gin
	r := gin.Default()

	route.RegisterSlotRoutes(r)

	r.Run(":8080")
}
