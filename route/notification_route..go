package route

import (
	"telemedicine-api/handler"

	"github.com/gin-gonic/gin"
)

func RegisterNotificationRoutes(r *gin.Engine) {
    r.GET("/api/notification", handler.GetNotifications)
    r.POST("/api/feedback", handler.PostFeedback)
    r.PATCH("/api/admin/verify-doctor/:id", handler.VerifyDoctor)
}