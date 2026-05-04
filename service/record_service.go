package service

import (
	"telemedicine-api/model"
	"telemedicine-api/repository"
)

type RecordService struct {
	Repo *repository.RecordRepository
}

func (s *RecordService) CreateRecord(rec *model.MedicalRecord) error {
	return s.Repo.CreateRecord(rec)
}

func (s *RecordService) GetPatientRecords(patientID int) ([]model.MedicalRecord, error) {
	return s.Repo.FindByPatientID(patientID)
}

func (s *RecordService) GetRecordByID(id int) (*model.MedicalRecord, error) {
	return s.Repo.FindByID(id)
}
