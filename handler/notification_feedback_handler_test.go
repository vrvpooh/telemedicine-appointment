package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"telemedicine-api/config"
	"telemedicine-api/model"
	"telemedicine-api/repository"
	"telemedicine-api/service"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupNFTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	
	config.ConnectDatabase()

	
	noteRepo := &repository.NotificationRepository{}
	noteSvc := &service.NotificationService{Repo: noteRepo}
	SetupNotificationHandler(noteSvc)

	feedRepo := &repository.FeedbackRepository{}
	feedSvc := &service.FeedbackService{Repo: feedRepo}
	SetupFeedbackHandler(feedSvc)

	
	r.GET("/api/notification", GetNotifications)
	r.POST("/api/feedback", PostFeedback)
	r.PATCH("/api/admin/verify-doctor/:id", VerifyDoctor)

	return r
}

func TestGetNotifications(t *testing.T) {
	router := setupNFTestRouter()
	
	t.Run("Get Notifications Success", func(t *testing.T) {
		
		req, _ := http.NewRequest("GET", "/api/notification?user_id=1", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})
}

func TestPostFeedback(t *testing.T) {
	router := setupNFTestRouter()

	t.Run("Post Feedback Success", func(t *testing.T) {
		input := model.Feedback{
			UserID:   1,
			DoctorID: 1,
			Rating:   5,
			Comment:  "Integration Test Feedback",
		}
		body, _ := json.Marshal(input)
		
		req, _ := http.NewRequest("POST", "/api/feedback", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		
		router.ServeHTTP(w, req)
		
		
		assert.True(t, w.Code == http.StatusOK || w.Code == http.StatusCreated)
	})
}

func TestVerifyDoctor(t *testing.T) {
	router := setupNFTestRouter()

	t.Run("Admin Verify Success", func(t *testing.T) {
		body := map[string]bool{"is_verified": true}
		jsonBody, _ := json.Marshal(body)

		
		req, _ := http.NewRequest("PATCH", "/api/admin/verify-doctor/1", bytes.NewBuffer(jsonBody))
		w := httptest.NewRecorder()
		
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
	})
}