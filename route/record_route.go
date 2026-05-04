package route

import (
	"telemedicine-api/handler"

	"github.com/gin-gonic/gin"
)

func RegisterRecordRoutes(r *gin.Engine) {
	recordGroup := r.Group("/api/records")
	{
		recordGroup.POST("", handler.CreateRecord)
		recordGroup.GET("/patient/me", handler.GetMyRecords)
		recordGroup.GET("/:id", handler.GetRecordByID)
	}
}
