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

	// Setup Appointment module
	appRepo := &repository.AppointmentRepository{}
	appService := &service.AppointmentService{Repo: appRepo, SlotRepo: slotRepo}
	handler.SetupAppointmentHandler(appService)

	// Setup Notification
	noteRepo := &repository.NotificationRepository{}
	noteService := &service.NotificationService{Repo: noteRepo}
	handler.SetupNotificationHandler(noteService)

	// Setup Feedback & Admin Verify
	feedRepo := &repository.FeedbackRepository{}
	feedService := &service.FeedbackService{Repo: feedRepo}
	handler.SetupFeedbackHandler(feedService)

	// Setup Record module
	recordRepo := &repository.RecordRepository{}
	recordService := &service.RecordService{Repo: recordRepo}
	handler.SetupRecordHandler(recordService)

	// Gin
	r := gin.Default()

	// Register routes
	route.RegisterSlotRoutes(r)
	route.RegisterDoctorRoutes(r)
	route.RegisterNotificationRoutes(r)
	route.RegisterRecordRoutes(r)
	route.RegisterAppointmentRoutes(r)

	r.Run(":8080")
}
