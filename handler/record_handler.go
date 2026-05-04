package handler

import (
	"net/http"
	"strconv"
	"telemedicine-api/model"
	"telemedicine-api/service"

	"github.com/gin-gonic/gin"
)

var recordService *service.RecordService

func SetupRecordHandler(s *service.RecordService) {
	recordService = s
}

func CreateRecord(c *gin.Context) {
	var rec model.MedicalRecord
	if err := c.ShouldBindJSON(&rec); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := recordService.CreateRecord(&rec); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create record"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Record created successfully", "data": rec})
}

func GetMyRecords(c *gin.Context) {
	// Mock PatientID ไปก่อนจนกว่าเพื่อนที่ทำ User/Auth จะส่ง JWT เสร็จ
	patientID := 101

	records, err := recordService.GetPatientRecords(patientID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch records"})
		return
	}
	if records == nil {
		records = []model.MedicalRecord{} // ป้องกันไม่ให้ Return เป็น null
	}
	c.JSON(http.StatusOK, gin.H{"data": records})
}

func GetRecordByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid record ID"})
		return
	}

	record, err := recordService.GetRecordByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	if record == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": record})
}
