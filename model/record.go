package model

type MedicalRecord struct {
	ID            int    `json:"id"`
	AppointmentID int    `json:"appointment_id"`
	PatientID     int    `json:"patient_id"`
	DoctorID      int    `json:"doctor_id"`
	Symptoms      string `json:"symptoms"`
	Diagnosis     string `json:"diagnosis"`
	Prescription  string `json:"prescription"`
	Notes         string `json:"notes"`
	CreatedAt     string `json:"created_at"`
}
