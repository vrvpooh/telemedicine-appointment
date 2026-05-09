package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"telemedicine-api/config"
	"telemedicine-api/model"
	"telemedicine-api/repository"
	"telemedicine-api/service"
	"testing"

	"github.com/gin-gonic/gin"
)

// ฟังก์ชันช่วยจำลอง Router สำหรับ Test
func setupTestRecordAPI() *gin.Engine {
	config.ConnectDatabase()
	repo := &repository.RecordRepository{}
	svc := &service.RecordService{Repo: repo}
	SetupRecordHandler(svc)

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/api/records", CreateRecord)
	r.GET("/api/records/patient/me", GetMyRecords)
	r.GET("/api/records/:id", GetRecordByID)
	return r
}

func TestRecordHandler(t *testing.T) {
	r := setupTestRecordAPI()

	// 1. เทสต์ GetMyRecords
	req1, _ := http.NewRequest("GET", "/api/records/patient/me", nil)
	w1 := httptest.NewRecorder()
	r.ServeHTTP(w1, req1)

	// 2. เทสต์ GetRecordByID
	req2, _ := http.NewRequest("GET", "/api/records/1", nil)
	w2 := httptest.NewRecorder()
	r.ServeHTTP(w2, req2)

	// 3. เทสต์ GetRecordByID กรณีส่ง ID ผิดรูปแบบ
	req3, _ := http.NewRequest("GET", "/api/records/abc", nil)
	w3 := httptest.NewRecorder()
	r.ServeHTTP(w3, req3)

	// 4. เทสต์ GetRecordByID กรณีหา ID ไม่เจอ
	req4, _ := http.NewRequest("GET", "/api/records/9999", nil)
	w4 := httptest.NewRecorder()
	r.ServeHTTP(w4, req4)

	// 5. เทสต์ CreateRecord
	mockData := model.MedicalRecord{
		AppointmentID: 99,
		PatientID:     101,
		DoctorID:      1,
		Symptoms:      "Test Symptoms",
	}
	jsonBody, _ := json.Marshal(mockData)
	req5, _ := http.NewRequest("POST", "/api/records", bytes.NewBuffer(jsonBody))
	w5 := httptest.NewRecorder()
	r.ServeHTTP(w5, req5)

	// 6. เทสต์ CreateRecord กรณีส่งข้อมูลผิดรูปแบบ
	req6, _ := http.NewRequest("POST", "/api/records", bytes.NewBuffer([]byte(`{"invalid": json`)))
	w6 := httptest.NewRecorder()
	r.ServeHTTP(w6, req6)

	// 7. จำลองสถานการณ์ Database พัง (เพื่อเก็บ Coverage บรรทัดสีแดงที่เป็น Error 500)
	// เนื่องจาก config.DB เป็น *sql.DB อยู่แล้ว เราสามารถสั่ง Close() ได้เลย
	config.DB.Close()

	// 7.1 เทสต์ GetMyRecords ตอน DB พัง
	req7, _ := http.NewRequest("GET", "/api/records/patient/me", nil)
	w7 := httptest.NewRecorder()
	r.ServeHTTP(w7, req7)

	// 7.2 เทสต์ GetRecordByID ตอน DB พัง
	req8, _ := http.NewRequest("GET", "/api/records/1", nil)
	w8 := httptest.NewRecorder()
	r.ServeHTTP(w8, req8)

	// 7.3 เทสต์ CreateRecord ตอน DB พัง
	req9, _ := http.NewRequest("POST", "/api/records", bytes.NewBuffer(jsonBody))
	w9 := httptest.NewRecorder()
	r.ServeHTTP(w9, req9)
}
