package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"telemedicine-api/model"

	"github.com/gin-gonic/gin"
)

// MockAppointmentService
type MockAppointmentService struct {
	BookAppointmentFunc func(patientID, slotID uint) (model.Appointment, error)
	GetZoomTokenFunc    func(id uint) (string, error)
	UpdateStatusFunc    func(id uint, status string) (model.Appointment, error)
}

func (m *MockAppointmentService) BookAppointment(patientID, slotID uint) (model.Appointment, error) {
	return m.BookAppointmentFunc(patientID, slotID)
}

func (m *MockAppointmentService) GetZoomToken(id uint) (string, error) {
	return m.GetZoomTokenFunc(id)
}

func (m *MockAppointmentService) UpdateStatus(id uint, status string) (model.Appointment, error) {
	return m.UpdateStatusFunc(id, status)
}

func TestBook_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockSvc := &MockAppointmentService{
		BookAppointmentFunc: func(patientID, slotID uint) (model.Appointment, error) {
			return model.Appointment{ID: 1, PatientID: patientID, SlotID: slotID, Status: "confirmed"}, nil
		},
	}
	SetupAppointmentHandler(mockSvc)

	r := gin.Default()
	r.POST("/appointments", Book)

	payload := map[string]uint{
		"patient_id": 1,
		"slot_id":    100,
	}
	body, _ := json.Marshal(payload)
	req, _ := http.NewRequest(http.MethodPost, "/appointments", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("expected status 201, got %d", w.Code)
	}

	var response model.Appointment
	json.Unmarshal(w.Body.Bytes(), &response)

	if response.ID != 1 {
		t.Errorf("expected appointment ID 1, got %d", response.ID)
	}
}

func TestBook_InvalidData(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	r.POST("/appointments", Book)

	payload := map[string]string{
		"patient_id": "invalid",
	}
	body, _ := json.Marshal(payload)
	req, _ := http.NewRequest(http.MethodPost, "/appointments", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected status 400, got %d", w.Code)
	}
}

func TestBook_ServiceError(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockSvc := &MockAppointmentService{
		BookAppointmentFunc: func(patientID, slotID uint) (model.Appointment, error) {
			return model.Appointment{}, errors.New("slot is already booked")
		},
	}
	SetupAppointmentHandler(mockSvc)

	r := gin.Default()
	r.POST("/appointments", Book)

	payload := map[string]uint{
		"patient_id": 1,
		"slot_id":    100,
	}
	body, _ := json.Marshal(payload)
	req, _ := http.NewRequest(http.MethodPost, "/appointments", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("expected status 500, got %d", w.Code)
	}

	var response map[string]string
	json.Unmarshal(w.Body.Bytes(), &response)

	if response["error"] != "slot is already booked" {
		t.Errorf("expected 'slot is already booked' error, got %s", response["error"])
	}
}

func TestGetToken_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockSvc := &MockAppointmentService{
		GetZoomTokenFunc: func(id uint) (string, error) {
			return "TOKEN123", nil
		},
	}
	SetupAppointmentHandler(mockSvc)

	r := gin.Default()
	r.GET("/appointments/:id/zoom-token", GetToken)

	req, _ := http.NewRequest(http.MethodGet, "/appointments/1/zoom-token", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}

	var response map[string]string
	json.Unmarshal(w.Body.Bytes(), &response)

	if response["zoom_token"] != "TOKEN123" {
		t.Errorf("expected TOKEN123, got %s", response["zoom_token"])
	}
}

func TestGetToken_InvalidID(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	r.GET("/appointments/:id/zoom-token", GetToken)

	req, _ := http.NewRequest(http.MethodGet, "/appointments/abc/zoom-token", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected status 400, got %d", w.Code)
	}
}

func TestGetToken_NotFound(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockSvc := &MockAppointmentService{
		GetZoomTokenFunc: func(id uint) (string, error) {
			return "", errors.New("not found")
		},
	}
	SetupAppointmentHandler(mockSvc)

	r := gin.Default()
	r.GET("/appointments/:id/zoom-token", GetToken)

	req, _ := http.NewRequest(http.MethodGet, "/appointments/999/zoom-token", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("expected status 404, got %d", w.Code)
	}
}

func TestUpdate_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockSvc := &MockAppointmentService{
		UpdateStatusFunc: func(id uint, status string) (model.Appointment, error) {
			return model.Appointment{ID: id, Status: status, UpdatedAt: "2026-05-09 15:30:00"}, nil
		},
	}
	SetupAppointmentHandler(mockSvc)

	r := gin.Default()
	r.PATCH("/appointments/:id/status", Update)

	payload := map[string]string{
		"status": "confirmed",
	}
	body, _ := json.Marshal(payload)
	req, _ := http.NewRequest(http.MethodPatch, "/appointments/1/status", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}

	var response model.Appointment
	json.Unmarshal(w.Body.Bytes(), &response)
	if response.Status != "confirmed" || response.UpdatedAt == "" {
		t.Errorf("expected status confirmed and non-empty updated_at, got status %s and updated_at %s", response.Status, response.UpdatedAt)
	}
}

func TestUpdate_InvalidID(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	r.PATCH("/appointments/:id/status", Update)

	req, _ := http.NewRequest(http.MethodPatch, "/appointments/abc/status", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected status 400, got %d", w.Code)
	}
}

func TestUpdate_MissingStatus(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	r.PATCH("/appointments/:id/status", Update)

	payload := map[string]string{}
	body, _ := json.Marshal(payload)
	req, _ := http.NewRequest(http.MethodPatch, "/appointments/1/status", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected status 400, got %d", w.Code)
	}
}

func TestUpdate_ServiceError(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockSvc := &MockAppointmentService{
		UpdateStatusFunc: func(id uint, status string) (model.Appointment, error) {
			return model.Appointment{}, errors.New("database error")
		},
	}
	SetupAppointmentHandler(mockSvc)

	r := gin.Default()
	r.PATCH("/appointments/:id/status", Update)

	payload := map[string]string{
		"status": "confirmed",
	}
	body, _ := json.Marshal(payload)
	req, _ := http.NewRequest(http.MethodPatch, "/appointments/1/status", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("expected status 500, got %d", w.Code)
	}
}
