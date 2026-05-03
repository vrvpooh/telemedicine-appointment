package handler

import (
	"net/http"

	"telemedicine-api/service"

	"github.com/gin-gonic/gin"
)

func GetDoctors(c *gin.Context) {
	doctors := service.GetDoctors()

	c.JSON(http.StatusOK, gin.H{
		"data": doctors,
	})
}

func GetDoctorByID(c *gin.Context) {
	id := c.Param("id")

	doctor, found := service.GetDoctorByID(id)
	if !found {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Doctor not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": doctor,
	})
}
