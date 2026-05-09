package repository

import (
	"database/sql"
	"telemedicine-api/config"
	"telemedicine-api/model"
	"time"
)

type RecordRepository struct{}

func (r *RecordRepository) CreateRecord(rec *model.MedicalRecord) error {
	query := `INSERT INTO medical_records (appointment_id, patient_id, doctor_id, symptoms, diagnosis, prescription, notes, created_at)
	          VALUES (?, ?, ?, ?, ?, ?, ?, ?)`
	now := time.Now().Format("2006-01-02 15:04:05")

	result, err := config.DB.Exec(query, rec.AppointmentID, rec.PatientID, rec.DoctorID, rec.Symptoms, rec.Diagnosis, rec.Prescription, rec.Notes, now)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err == nil {
		rec.ID = int(id)
		rec.CreatedAt = now
	}
	return nil
}

func (r *RecordRepository) FindByPatientID(patientID int) ([]model.MedicalRecord, error) {
	query := `SELECT id, appointment_id, patient_id, doctor_id, symptoms, diagnosis, prescription, notes, created_at FROM medical_records WHERE patient_id = ?`
	rows, err := config.DB.Query(query, patientID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var records []model.MedicalRecord
	for rows.Next() {
		var rec model.MedicalRecord
		if err := rows.Scan(&rec.ID, &rec.AppointmentID, &rec.PatientID, &rec.DoctorID, &rec.Symptoms, &rec.Diagnosis, &rec.Prescription, &rec.Notes, &rec.CreatedAt); err != nil {
			return nil, err
		}
		records = append(records, rec)
	}
	return records, nil
}

func (r *RecordRepository) FindByID(id int) (*model.MedicalRecord, error) {
	query := `SELECT id, appointment_id, patient_id, doctor_id, symptoms, diagnosis, prescription, notes, created_at FROM medical_records WHERE id = ?`
	row := config.DB.QueryRow(query, id)

	var rec model.MedicalRecord
	err := row.Scan(&rec.ID, &rec.AppointmentID, &rec.PatientID, &rec.DoctorID, &rec.Symptoms, &rec.Diagnosis, &rec.Prescription, &rec.Notes, &rec.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // หาไม่เจอ
		}
		return nil, err
	}
	return &rec, nil
}
