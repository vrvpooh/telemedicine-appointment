package route

import (
	"telemedicine-api/handler"

	"github.com/gin-gonic/gin"
)

func RegisterAppointmentRoutes(r *gin.Engine) {
	g := r.Group("/api/appointments")
	{
		g.POST("", handler.Book)
		g.GET("/:id/zoom-token", handler.GetToken)
		g.PATCH("/:id/status", handler.Update)
	}
}
