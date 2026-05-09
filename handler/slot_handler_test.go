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

// MockSlotService
type MockSlotService struct {
	CreateSlotFunc        func(doctorID int64, slot model.Slot) (model.Slot, error)
	GetAvailableSlotsFunc func(doctorID string) ([]model.Slot, error)
	DeleteSlotFunc        func(id string) (int64, error)
}

func (m *MockSlotService) CreateSlot(doctorID int64, slot model.Slot) (model.Slot, error) {
	return m.CreateSlotFunc(doctorID, slot)
}
func (m *MockSlotService) GetAvailableSlots(doctorID string) ([]model.Slot, error) {
	return m.GetAvailableSlotsFunc(doctorID)
}
func (m *MockSlotService) DeleteSlot(id string) (int64, error) {
	return m.DeleteSlotFunc(id)
}

// Create Slot Test

func TestCreateSlot_HandlerSuccess(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockSvc := &MockSlotService{
		CreateSlotFunc: func(doctorID int64, slot model.Slot) (model.Slot, error) {
			slot.ID = 1
			slot.DoctorID = doctorID
			return slot, nil
		},
	}
	SetupSlotHandler(mockSvc)

	r := gin.Default()
	r.POST("/api/doctor/:id/slots", CreateSlot)

	body := []byte(`{"start_time":"2026-05-15 09:00","end_time":"2026-05-15 10:00"}`)
	req, _ := http.NewRequest(http.MethodPost, "/api/doctor/1/slots", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("expected status 201, got %d", w.Code)
	}
}

func TestCreateSlot_InvalidDoctorID(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	r.POST("/api/doctor/:id/slots", CreateSlot)

	body := []byte(`{"start_time":"2026-05-15 09:00","end_time":"2026-05-15 10:00"}`)
	req, _ := http.NewRequest(http.MethodPost, "/api/doctor/abc/slots", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected status 400, got %d", w.Code)
	}
}

func TestCreateSlot_InvalidJSON(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	r.POST("/api/doctor/:id/slots", CreateSlot)

	body := []byte(`{invalid json`)
	req, _ := http.NewRequest(http.MethodPost, "/api/doctor/1/slots", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected status 400, got %d", w.Code)
	}
}

func TestCreateSlot_ServiceError(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockSvc := &MockSlotService{
		CreateSlotFunc: func(doctorID int64, slot model.Slot) (model.Slot, error) {
			return slot, errors.New("invalid time format")
		},
	}
	SetupSlotHandler(mockSvc)

	r := gin.Default()
	r.POST("/api/doctor/:id/slots", CreateSlot)

	body := []byte(`{"start_time":"bad","end_time":"bad"}`)
	req, _ := http.NewRequest(http.MethodPost, "/api/doctor/1/slots", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected status 400, got %d", w.Code)
	}
}

// GetAvailableSlots Test

func TestGetAvailableSlots_HandlerSuccess(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockSvc := &MockSlotService{
		GetAvailableSlotsFunc: func(doctorID string) ([]model.Slot, error) {
			return []model.Slot{
				{ID: 1, DoctorID: 1, StartTime: "2026-05-15 09:00", EndTime: "2026-05-15 10:00"},
			}, nil
		},
	}
	SetupSlotHandler(mockSvc)

	r := gin.Default()
	r.GET("/api/doctor/:id/slots", GetAvailableSlots)

	req, _ := http.NewRequest(http.MethodGet, "/api/doctor/1/slots", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}

	var resp map[string]interface{}
	_ = json.Unmarshal(w.Body.Bytes(), &resp)
	if resp["count"].(float64) != 1 {
		t.Errorf("expected count 1, got %v", resp["count"])
	}
}

func TestGetAvailableSlots_Empty(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockSvc := &MockSlotService{
		GetAvailableSlotsFunc: func(doctorID string) ([]model.Slot, error) {
			return []model.Slot{}, nil
		},
	}
	SetupSlotHandler(mockSvc)

	r := gin.Default()
	r.GET("/api/doctor/:id/slots", GetAvailableSlots)

	req, _ := http.NewRequest(http.MethodGet, "/api/doctor/999/slots", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}
}

func TestGetAvailableSlots_ServiceError(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockSvc := &MockSlotService{
		GetAvailableSlotsFunc: func(doctorID string) ([]model.Slot, error) {
			return nil, errors.New("db error")
		},
	}
	SetupSlotHandler(mockSvc)

	r := gin.Default()
	r.GET("/api/doctor/:id/slots", GetAvailableSlots)

	req, _ := http.NewRequest(http.MethodGet, "/api/doctor/1/slots", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("expected status 500, got %d", w.Code)
	}
}

// DeleteSlot Test

func TestDeleteSlot_HandlerSuccess(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockSvc := &MockSlotService{
		DeleteSlotFunc: func(id string) (int64, error) {
			return 1, nil
		},
	}
	SetupSlotHandler(mockSvc)

	r := gin.Default()
	r.DELETE("/api/slots/:id", DeleteSlot)

	req, _ := http.NewRequest(http.MethodDelete, "/api/slots/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}
}

func TestDeleteSlot_NotFound(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockSvc := &MockSlotService{
		DeleteSlotFunc: func(id string) (int64, error) {
			return 0, nil
		},
	}
	SetupSlotHandler(mockSvc)

	r := gin.Default()
	r.DELETE("/api/slots/:id", DeleteSlot)

	req, _ := http.NewRequest(http.MethodDelete, "/api/slots/9999", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("expected status 404, got %d", w.Code)
	}
}

func TestDeleteSlot_ServiceError(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockSvc := &MockSlotService{
		DeleteSlotFunc: func(id string) (int64, error) {
			return 0, errors.New("db error")
		},
	}
	SetupSlotHandler(mockSvc)

	r := gin.Default()
	r.DELETE("/api/slots/:id", DeleteSlot)

	req, _ := http.NewRequest(http.MethodDelete, "/api/slots/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("expected status 500, got %d", w.Code)
	}
}
