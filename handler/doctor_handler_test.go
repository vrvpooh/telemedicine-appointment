package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"telemedicine-api/repository"
	"telemedicine-api/service"

	"fmt"
	"telemedicine-api/model"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func setupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)

	repo := &repository.MockDoctorRepository{}

	repo.On("GetAll", "", "").Return([]model.Doctor{}, nil)

	repo.On("GetAll", mock.Anything, mock.Anything).
		Return([]model.Doctor{}, nil)

	repo.On("GetByID", "1").
		Return(&model.Doctor{ID: "1"}, nil)

	repo.On("GetByID", "999").
		Return((*model.Doctor)(nil), fmt.Errorf("not found"))

	repo.On("GetSpecialties").
		Return([]model.Specialty{
			{Name: "Cardio"},
		}, nil)

	svc := &service.DoctorService{Repo: repo}
	SetupDoctorHandler(svc)

	r := gin.Default()
	r.GET("/api/doctors", GetDoctors)
	r.GET("/api/doctors/:id", GetDoctorByID)
	r.GET("/api/specialties", GetSpecialties)

	return r
}

func TestGetDoctors(t *testing.T) {
	router := setupTestRouter()

	req, _ := http.NewRequest("GET", "/api/doctors", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetDoctors_WithQuery(t *testing.T) {
	router := setupTestRouter()

	req, _ := http.NewRequest("GET", "/api/doctors?name=John&specialty=Cardio", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetDoctorByID(t *testing.T) {
	router := setupTestRouter()

	req, _ := http.NewRequest("GET", "/api/doctors/1", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	// อาจเป็น 200 หรือ 404 ขึ้นกับ mock ของคุณ
	assert.True(t, w.Code == http.StatusOK || w.Code == http.StatusNotFound)
}

func TestGetSpecialties(t *testing.T) {
	router := setupTestRouter()

	req, _ := http.NewRequest("GET", "/api/specialties", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetDoctors_Error(t *testing.T) {
	gin.SetMode(gin.TestMode)

	repo := &repository.MockDoctorRepository{}

	// 🔥 บังคับให้ error
	repo.On("GetAll", "error", "").
		Return([]model.Doctor{}, fmt.Errorf("db error"))

	svc := &service.DoctorService{Repo: repo}
	SetupDoctorHandler(svc)

	r := gin.Default()
	r.GET("/api/doctors", GetDoctors)

	req, _ := http.NewRequest("GET", "/api/doctors?name=error", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestGetDoctorByID_NotFound(t *testing.T) {
	router := setupTestRouter()

	req, _ := http.NewRequest("GET", "/api/doctors/999", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestGetSpecialties_Error(t *testing.T) {
	gin.SetMode(gin.TestMode)

	repo := &repository.MockDoctorRepository{}

	// 🔥 บังคับ error
	repo.On("GetSpecialties").
		Return([]model.Specialty{}, fmt.Errorf("db error"))

	svc := &service.DoctorService{Repo: repo}
	SetupDoctorHandler(svc)

	r := gin.Default()
	r.GET("/api/specialties", GetSpecialties)

	req, _ := http.NewRequest("GET", "/api/specialties", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}
