package repository

import (
	"telemedicine-api/config"
	"telemedicine-api/model"
)

type AppointmentRepository struct{}

func (r *AppointmentRepository) Create(app *model.Appointment) error {
	query := `INSERT INTO appointments (patient_id, slot_id, status, zoom_token) VALUES (?, ?, ?, ?)`
	_, err := config.DB.Exec(query, app.PatientID, app.SlotID, app.Status, app.ZoomToken)
	return err
}

func (r *AppointmentRepository) GetByID(id uint) (model.Appointment, error) {
	var app model.Appointment
	query := `SELECT id, patient_id, slot_id, status, zoom_token FROM appointments WHERE id = ?`
	err := config.DB.QueryRow(query, id).Scan(&app.ID, &app.PatientID, &app.SlotID, &app.Status, &app.ZoomToken)
	return app, err
}

func (r *AppointmentRepository) Update(app model.Appointment) error {
	query := `UPDATE appointments SET status = ? WHERE id = ?`
	_, err := config.DB.Exec(query, app.Status, app.ID)
	return err
}
