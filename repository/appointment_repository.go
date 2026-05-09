package repository

import (
	"database/sql"
	"errors"
	"telemedicine-api/config"
	"telemedicine-api/model"
	"time"
)

type IAppointmentRepository interface {
	Create(app *model.Appointment) error
	GetByID(id uint) (model.Appointment, error)
	Update(app model.Appointment) error
}

type AppointmentRepository struct{}

func (r *AppointmentRepository) Create(app *model.Appointment) error {
	now := time.Now().Format("2006-01-02 15:04:05")
	app.CreatedAt = now
	app.UpdatedAt = now

	query := `INSERT INTO appointments (patient_id, slot_id, status, zoom_token, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)`
	result, err := config.DB.Exec(query, app.PatientID, app.SlotID, app.Status, app.ZoomToken, app.CreatedAt, app.UpdatedAt)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err == nil {
		app.ID = uint(id)
	}
	return nil
}

func (r *AppointmentRepository) GetByID(id uint) (model.Appointment, error) {
	var app model.Appointment
	var createdAt, updatedAt string
	query := `SELECT id, patient_id, slot_id, status, zoom_token, created_at, updated_at FROM appointments WHERE id = ?`
	err := config.DB.QueryRow(query, id).Scan(&app.ID, &app.PatientID, &app.SlotID, &app.Status, &app.ZoomToken, &createdAt, &updatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return app, errors.New("appointment not found")
		}
		return app, err
	}

	// พยายามแปลง Format เวลาให้สวยงาม (รองรับทั้ง RFC3339 และ format อื่นๆ)
	app.CreatedAt = formatTimeStr(createdAt)
	app.UpdatedAt = formatTimeStr(updatedAt)

	return app, nil
}

// formatTimeStr ช่วยแปลง string เวลาจาก DB ให้เป็น format "2006-01-02 15:04:05"
func formatTimeStr(tStr string) string {
	layouts := []string{
		"2006-01-02 15:04:05",
		time.RFC3339,
		"2006-01-02T15:04:05Z07:00",
		"2006-01-02 15:04:05.999999999-07:00",
	}

	var t time.Time
	var err error
	for _, layout := range layouts {
		t, err = time.Parse(layout, tStr)
		if err == nil {
			return t.Format("2006-01-02 15:04:05")
		}
	}
	// ถ้าแปลงไม่ได้เลย ให้ส่งค่าเดิมกลับไป
	return tStr
}

func (r *AppointmentRepository) Update(app model.Appointment) error {
	app.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	query := `UPDATE appointments SET status = ?, updated_at = ? WHERE id = ?`
	_, err := config.DB.Exec(query, app.Status, app.UpdatedAt, app.ID)
	return err
}
