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
