package handler

import (
	"net/http"
	"telemedicine-api/model"
	"telemedicine-api/service"

	"github.com/gin-gonic/gin"
)

type FeedbackHandler struct {
    Service *service.FeedbackService
}

var feedbackHandler *FeedbackHandler

func SetupFeedbackHandler(s *service.FeedbackService) {
    feedbackHandler = &FeedbackHandler{Service: s}
}

func PostFeedback(c *gin.Context) {
    var req model.Feedback
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := feedbackHandler.Service.SubmitFeedback(req); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, gin.H{"message": "Feedback submitted"})
}

func VerifyDoctor(c *gin.Context) {
    id := c.Param("id")
    var body struct {
        IsVerified bool `json:"is_verified"`
    }
    c.ShouldBindJSON(&body)

    err := feedbackHandler.Service.VerifyDoctorStatus(id, body.IsVerified)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Doctor status updated"})
}