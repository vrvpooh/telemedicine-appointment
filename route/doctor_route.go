package route

import (
	"telemedicine-api/handler"

	"github.com/gin-gonic/gin"
)

func RegisterDoctorRoutes(r *gin.Engine) {
	api := r.Group("/api")

	api.GET("/doctors", handler.GetDoctors)
	api.GET("/doctors/:id", handler.GetDoctorByID)
}
