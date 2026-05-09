package handler

import (
	"net/http"
	"strconv"
	"telemedicine-api/service"

	"github.com/gin-gonic/gin"
)

var appSvc service.IAppointmentService

func SetupAppointmentHandler(s service.IAppointmentService) {
	appSvc = s
}

func Book(c *gin.Context) {
	var req struct {
		PatientID uint `json:"patient_id" binding:"required"`
		SlotID    uint `json:"slot_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	app, err := appSvc.BookAppointment(req.PatientID, req.SlotID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, app)
}

func GetToken(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid appointment ID"})
		return
	}

	token, err := appSvc.GetZoomToken(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Appointment not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"zoom_token": token})
}

func Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid appointment ID"})
		return
	}

	var req struct {
		Status string `json:"status" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Status is required"})
		return
	}

	app, err := appSvc.UpdateStatus(uint(id), req.Status)
	if err != nil {
		if err.Error() == "appointment not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, app)
}
