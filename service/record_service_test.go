package service

import (
	"telemedicine-api/config"
	"telemedicine-api/model"
	"telemedicine-api/repository"
	"testing"
)

func TestRecordService(t *testing.T) {
	// เตรียม Database และ Inject เข้า Service
	config.ConnectDatabase()
	repo := &repository.RecordRepository{}
	svc := &RecordService{Repo: repo}

	// 1. เทสต์ GetPatientRecords
	_, err := svc.GetPatientRecords(101)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// 2. เทสต์ GetRecordByID (ดึง ID = 1 ซึ่งเรามีใน Mock Data)
	_, err = svc.GetRecordByID(1)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// 3. เทสต์ CreateRecord
	newRecord := &model.MedicalRecord{
		AppointmentID: 88,
		PatientID:     101,
		DoctorID:      2,
		Symptoms:      "Service Test",
	}
	err = svc.CreateRecord(newRecord)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}
