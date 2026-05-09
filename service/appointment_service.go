package service

import (
	"database/sql"
	"errors"
	"strconv"
	"telemedicine-api/model"
	"telemedicine-api/repository"
	"time"
)

type IAppointmentService interface {
	BookAppointment(patientID, slotID uint) (model.Appointment, error)
	GetZoomToken(id uint) (string, error)
	UpdateStatus(id uint, status string) (model.Appointment, error)
}

type AppointmentService struct {
	DB       *sql.DB
	Repo     repository.IAppointmentRepository
	SlotRepo repository.ISlotRepository
}

func (s *AppointmentService) BookAppointment(patientID, slotID uint) (model.Appointment, error) {
	// เริ่ม Transaction เพื่อความปลอดภัย
	tx, err := s.DB.Begin()
	if err != nil {
		return model.Appointment{}, err
	}
	defer tx.Rollback()

	// 1. ตรวจสอบ Slot (Lock เพื่อป้องกันคนอื่นแก้ไขพร้อมกัน)
	var isBooked bool
	err = tx.QueryRow("SELECT is_booked FROM slots WHERE id = ?", slotID).Scan(&isBooked)
	if err != nil {
		return model.Appointment{}, errors.New("ไม่พบ slot ที่ระบุ")
	}

	if isBooked {
		return model.Appointment{}, errors.New("slot นี้ถูกจองไปแล้ว")
	}

	// 2. สร้าง Appointment
	app := model.Appointment{
		PatientID: patientID,
		SlotID:    slotID,
		Status:    "confirmed",
		ZoomToken: "ZOOM-" + strconv.Itoa(int(slotID)),
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
		UpdatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}

	query := `INSERT INTO appointments (patient_id, slot_id, status, zoom_token, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)`
	result, err := tx.Exec(query, app.PatientID, app.SlotID, app.Status, app.ZoomToken, app.CreatedAt, app.UpdatedAt)
	if err != nil {
		// ถ้า Error เพราะ UNIQUE constraint ทำงาน (slot_id ซ้ำ)
		return model.Appointment{}, errors.New("slot นี้ถูกจองไปแล้ว (Database constraint)")
	}

	id, _ := result.LastInsertId()
	app.ID = uint(id)

	// 3. อัปเดต Slot Status
	_, err = tx.Exec("UPDATE slots SET is_booked = 1 WHERE id = ?", slotID)
	if err != nil {
		return model.Appointment{}, err
	}

	// Commit transaction
	if err := tx.Commit(); err != nil {
		return model.Appointment{}, err
	}

	return app, nil
}

func (s *AppointmentService) GetZoomToken(id uint) (string, error) {
	app, err := s.Repo.GetByID(id)
	return app.ZoomToken, err
}

func (s *AppointmentService) UpdateStatus(id uint, status string) (model.Appointment, error) {
	app, err := s.Repo.GetByID(id)
	if err != nil {
		return model.Appointment{}, err
	}
	app.Status = status
	app.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")

	if err := s.Repo.Update(app); err != nil {
		return model.Appointment{}, err
	}

	return app, nil
}
