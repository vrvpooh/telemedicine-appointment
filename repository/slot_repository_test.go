package repository

import (
	"database/sql"
	"testing"

	"telemedicine-api/config"
	"telemedicine-api/model"

	_ "github.com/mattn/go-sqlite3"
)

// setupTestDB
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("failed to open in-memory db: %v", err)
	}

	schema := `
	CREATE TABLE slots (
		id         INTEGER PRIMARY KEY AUTOINCREMENT,
		doctor_id  INTEGER NOT NULL,
		start_time TEXT    NOT NULL,
		end_time   TEXT    NOT NULL,
		is_booked  INTEGER NOT NULL DEFAULT 0
	);
	`
	if _, err := db.Exec(schema); err != nil {
		t.Fatalf("failed to create schema: %v", err)
	}

	config.DB = db
	return db
}

func TestCreate_Success(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	repo := &SlotRepository{}
	slot := model.Slot{
		DoctorID:  1,
		StartTime: "2026-05-15 09:00",
		EndTime:   "2026-05-15 10:00",
	}

	created, err := repo.Create(slot)

	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if created.ID == 0 {
		t.Error("expected ID to be auto-generated, got 0")
	}
	if created.IsBooked != false {
		t.Errorf("expected IsBooked false, got %v", created.IsBooked)
	}
}

func TestCreate_DBError(t *testing.T) {
	db := setupTestDB(t)
	db.Close()

	repo := &SlotRepository{}
	slot := model.Slot{DoctorID: 1, StartTime: "x", EndTime: "y"}

	_, err := repo.Create(slot)
	if err == nil {
		t.Error("expected error when DB is closed, got nil")
	}
}

func TestGetAvailableByDoctorID_Success(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	// insert 3 slots:
	// 1) doctor 1 ว่าง 1 case | 2) doctor 1 ถูกจอง 1 case | 3) doctor 2
	_, _ = db.Exec(`INSERT INTO slots (doctor_id, start_time, end_time, is_booked) VALUES
		(1, '2026-05-15 09:00', '2026-05-15 10:00', 0),
		(1, '2026-05-15 10:00', '2026-05-15 11:00', 1),
		(2, '2026-05-15 14:00', '2026-05-15 15:00', 0)`)

	repo := &SlotRepository{}
	slots, err := repo.GetAvailableByDoctorID("1")

	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if len(slots) != 1 {
		t.Errorf("expected 1 available slot for doctor 1, got %d", len(slots))
	}
	if len(slots) > 0 && slots[0].IsBooked {
		t.Error("expected only available slots (is_booked=0)")
	}
}

func TestGetAvailableByDoctorID_Empty(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	repo := &SlotRepository{}
	slots, err := repo.GetAvailableByDoctorID("999")

	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if len(slots) != 0 {
		t.Errorf("expected 0 slots, got %d", len(slots))
	}
}

func TestGetAvailableByDoctorID_DBError(t *testing.T) {
	db := setupTestDB(t)
	db.Close()

	repo := &SlotRepository{}
	_, err := repo.GetAvailableByDoctorID("1")
	if err == nil {
		t.Error("expected error when DB is closed, got nil")
	}
}

func TestDeleteByID_Success(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	_, _ = db.Exec(`INSERT INTO slots (id, doctor_id, start_time, end_time, is_booked)
		VALUES (1, 1, '2026-05-15 09:00', '2026-05-15 10:00', 0)`)

	repo := &SlotRepository{}
	rows, err := repo.DeleteByID("1")

	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if rows != 1 {
		t.Errorf("expected 1 row affected, got %d", rows)
	}
}

func TestDeleteByID_NotFound(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	repo := &SlotRepository{}
	rows, err := repo.DeleteByID("9999")

	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if rows != 0 {
		t.Errorf("expected 0 rows affected, got %d", rows)
	}
}

func TestDeleteByID_DBError(t *testing.T) {
	db := setupTestDB(t)
	db.Close()

	repo := &SlotRepository{}
	_, err := repo.DeleteByID("1")
	if err == nil {
		t.Error("expected error when DB is closed, got nil")
	}
}

func TestGetByID_Success(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	_, _ = db.Exec(`INSERT INTO slots (id, doctor_id, start_time, end_time, is_booked)
		VALUES (5, 2, '2026-05-15 09:00', '2026-05-15 10:00', 0)`)

	repo := &SlotRepository{}
	slot, err := repo.GetByID(5)

	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if slot.ID != 5 {
		t.Errorf("expected ID 5, got %d", slot.ID)
	}
	if slot.DoctorID != 2 {
		t.Errorf("expected DoctorID 2, got %d", slot.DoctorID)
	}
}

func TestGetByID_NotFound(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	repo := &SlotRepository{}
	_, err := repo.GetByID(9999)

	if err == nil {
		t.Error("expected error for non-existent slot, got nil")
	}
}

func TestUpdateIsBooked_Success(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	_, _ = db.Exec(`INSERT INTO slots (id, doctor_id, start_time, end_time, is_booked)
		VALUES (1, 1, '2026-05-15 09:00', '2026-05-15 10:00', 0)`)

	repo := &SlotRepository{}
	err := repo.UpdateIsBooked(1, true)

	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	// verify DB state
	var isBooked bool
	_ = db.QueryRow("SELECT is_booked FROM slots WHERE id = 1").Scan(&isBooked)
	if !isBooked {
		t.Error("expected is_booked = true after update")
	}
}

func TestUpdateIsBooked_DBError(t *testing.T) {
	db := setupTestDB(t)
	db.Close()

	repo := &SlotRepository{}
	err := repo.UpdateIsBooked(1, true)
	if err == nil {
		t.Error("expected error when DB is closed, got nil")
	}
}
