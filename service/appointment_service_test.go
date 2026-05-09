package service

import (
	"database/sql"
	"errors"
	"testing"

	"telemedicine-api/model"

	_ "github.com/mattn/go-sqlite3"
)

func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("failed to open in-memory db: %v", err)
	}

	schema := `
	CREATE TABLE slots (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		doctor_id INTEGER,
		start_time TEXT,
		end_time TEXT,
		is_booked INTEGER DEFAULT 0
	);
	CREATE TABLE appointments (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		patient_id INTEGER,
		slot_id INTEGER UNIQUE,
		status TEXT,
		zoom_token TEXT,
		created_at TEXT,
		updated_at TEXT
	);
	`
	if _, err := db.Exec(schema); err != nil {
		t.Fatalf("failed to create schema: %v", err)
	}
	return db
}

// MockAppointmentRepository
type MockAppointmentRepository struct {
	CreateFunc  func(app *model.Appointment) error
	GetByIDFunc func(id uint) (model.Appointment, error)
	UpdateFunc  func(app model.Appointment) error
}

func (m *MockAppointmentRepository) Create(app *model.Appointment) error {
	return m.CreateFunc(app)
}

func (m *MockAppointmentRepository) GetByID(id uint) (model.Appointment, error) {
	return m.GetByIDFunc(id)
}

func (m *MockAppointmentRepository) Update(app model.Appointment) error {
	return m.UpdateFunc(app)
}

// MockSlotRepository
type MockSlotRepository struct {
	CreateFunc                 func(slot model.Slot) (model.Slot, error)
	GetAvailableByDoctorIDFunc func(doctorID string) ([]model.Slot, error)
	DeleteByIDFunc             func(id string) (int64, error)
	GetByIDFunc                func(id uint) (model.Slot, error)
	UpdateIsBookedFunc         func(id uint, isBooked bool) error
}

func (m *MockSlotRepository) Create(slot model.Slot) (model.Slot, error) {
	return m.CreateFunc(slot)
}

func (m *MockSlotRepository) GetAvailableByDoctorID(doctorID string) ([]model.Slot, error) {
	return m.GetAvailableByDoctorIDFunc(doctorID)
}

func (m *MockSlotRepository) DeleteByID(id string) (int64, error) {
	return m.DeleteByIDFunc(id)
}

func (m *MockSlotRepository) GetByID(id uint) (model.Slot, error) {
	return m.GetByIDFunc(id)
}

func (m *MockSlotRepository) UpdateIsBooked(id uint, isBooked bool) error {
	return m.UpdateIsBookedFunc(id, isBooked)
}

func TestBookAppointment_Success(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	// Insert a mock slot
	_, _ = db.Exec("INSERT INTO slots (id, doctor_id, start_time, end_time, is_booked) VALUES (100, 1, '2026-05-10 10:00', '2026-05-10 11:00', 0)")

	svc := &AppointmentService{
		DB: db,
	}

	app, err := svc.BookAppointment(1, 100)

	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if app.PatientID != 1 {
		t.Errorf("expected patient ID 1, got %d", app.PatientID)
	}
	if app.SlotID != 100 {
		t.Errorf("expected slot ID 100, got %d", app.SlotID)
	}
	if app.Status != "confirmed" {
		t.Errorf("expected status confirmed, got %s", app.Status)
	}

	// Verify DB state
	var isBooked bool
	_ = db.QueryRow("SELECT is_booked FROM slots WHERE id = 100").Scan(&isBooked)
	if !isBooked {
		t.Error("expected slot to be marked as booked in DB")
	}
}

func TestBookAppointment_AlreadyBooked(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	_, _ = db.Exec("INSERT INTO slots (id, doctor_id, start_time, end_time, is_booked) VALUES (100, 1, '10:00', '11:00', 1)")

	svc := &AppointmentService{
		DB: db,
	}

	_, err := svc.BookAppointment(1, 100)

	if err == nil {
		t.Error("expected error for already booked slot, got nil")
	}
	if err.Error() != "slot นี้ถูกจองไปแล้ว" {
		t.Errorf("expected 'slot นี้ถูกจองไปแล้ว' error, got %v", err)
	}
}

func TestBookAppointment_SlotNotFound(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	svc := &AppointmentService{
		DB: db,
	}

	_, err := svc.BookAppointment(1, 100)

	if err == nil {
		t.Error("expected error for non-existent slot, got nil")
	}
	if err.Error() != "ไม่พบ slot ที่ระบุ" {
		t.Errorf("expected 'ไม่พบ slot ที่ระบุ' error, got %v", err)
	}
}

func TestBookAppointment_CreateError(t *testing.T) {
	db := setupTestDB(t)
	// ปิด DB ทันทีเพื่อให้เกิด Error ตอน Exec
	db.Close()

	svc := &AppointmentService{
		DB: db,
	}

	_, err := svc.BookAppointment(1, 100)

	if err == nil {
		t.Error("expected error when creation fails, got nil")
	}
}

func TestGetZoomToken_Success(t *testing.T) {
	mockRepo := &MockAppointmentRepository{
		GetByIDFunc: func(id uint) (model.Appointment, error) {
			return model.Appointment{ID: id, ZoomToken: "TOKEN123"}, nil
		},
	}

	svc := &AppointmentService{
		Repo: mockRepo,
	}

	token, err := svc.GetZoomToken(1)

	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if token != "TOKEN123" {
		t.Errorf("expected TOKEN123, got %s", token)
	}
}

func TestGetZoomToken_Error(t *testing.T) {
	mockRepo := &MockAppointmentRepository{
		GetByIDFunc: func(id uint) (model.Appointment, error) {
			return model.Appointment{}, errors.New("not found")
		},
	}

	svc := &AppointmentService{
		Repo: mockRepo,
	}

	_, err := svc.GetZoomToken(1)

	if err == nil {
		t.Error("expected error, got nil")
	}
}

func TestUpdateStatus_Success(t *testing.T) {
	updateCalled := false
	mockRepo := &MockAppointmentRepository{
		GetByIDFunc: func(id uint) (model.Appointment, error) {
			return model.Appointment{ID: id, Status: "pending"}, nil
		},
		UpdateFunc: func(app model.Appointment) error {
			if app.ID == 1 && app.Status == "cancelled" {
				updateCalled = true
			}
			return nil
		},
	}

	svc := &AppointmentService{
		Repo: mockRepo,
	}

	app, err := svc.UpdateStatus(1, "cancelled")

	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if !updateCalled {
		t.Error("expected Update to be called with correct status")
	}
	if app.Status != "cancelled" {
		t.Errorf("expected status cancelled, got %s", app.Status)
	}
}

func TestUpdateStatus_GetError(t *testing.T) {
	mockRepo := &MockAppointmentRepository{
		GetByIDFunc: func(id uint) (model.Appointment, error) {
			return model.Appointment{}, errors.New("not found")
		},
	}

	svc := &AppointmentService{
		Repo: mockRepo,
	}

	_, err := svc.UpdateStatus(1, "cancelled")

	if err == nil {
		t.Error("expected error when GetByID fails, got nil")
	}
}

func TestUpdateStatus_UpdateError(t *testing.T) {
	mockRepo := &MockAppointmentRepository{
		GetByIDFunc: func(id uint) (model.Appointment, error) {
			return model.Appointment{ID: id, Status: "pending"}, nil
		},
		UpdateFunc: func(app model.Appointment) error {
			return errors.New("update failed")
		},
	}

	svc := &AppointmentService{
		Repo: mockRepo,
	}

	_, err := svc.UpdateStatus(1, "cancelled")

	if err == nil {
		t.Error("expected error when Update fails, got nil")
	}
}

