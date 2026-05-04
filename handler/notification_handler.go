package handler

import (
	"net/http"
	"telemedicine-api/service"

	"github.com/gin-gonic/gin"
)

type NotificationHandler struct {
    Service *service.NotificationService
}

var notificationHandler *NotificationHandler

func SetupNotificationHandler(s *service.NotificationService) {
    notificationHandler = &NotificationHandler{Service: s}
}

func GetNotifications(c *gin.Context) {
    // สมมติ userID = 1 (ในอนาคตดึงจาก Auth Token)
    data, err := notificationHandler.Service.GetUserNotifications(1)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"data": data})
}