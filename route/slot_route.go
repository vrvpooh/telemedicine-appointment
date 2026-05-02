package route

import (
	"telemedicine-api/handler"

	"github.com/gin-gonic/gin"
)

func RegisterSlotRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.POST("/doctor/:id/slots", handler.CreateSlot)
		api.GET("/doctor/:id/slots", handler.GetAvailableSlots)
		api.DELETE("/slots/:id", handler.DeleteSlot)
	}
}
