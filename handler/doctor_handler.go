package handler

import (
	"net/http"

	"telemedicine-api/service"

	"github.com/gin-gonic/gin"
)

type DoctorHandler struct {
	Service *service.DoctorService
}

var doctorHandler *DoctorHandler

// setup handler (เหมือน slot)
func SetupDoctorHandler(s *service.DoctorService) {
	doctorHandler = &DoctorHandler{Service: s}
}

// GET /api/doctors
func GetDoctors(c *gin.Context) {
	doctors, err := doctorHandler.Service.GetDoctors()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": doctors})
}

// GET /api/doctors/:id
func GetDoctorByID(c *gin.Context) {
	id := c.Param("id")

	doctor, err := doctorHandler.Service.GetDoctorByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Doctor not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": doctor})
}

// GET /api/specialties
func GetSpecialties(c *gin.Context) {
	specialties, err := doctorHandler.Service.GetSpecialties()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": specialties})
}
