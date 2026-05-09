package service

import (
	"errors"
	"testing"

	"telemedicine-api/model"
)

// MockSlotRepoForService
type MockSlotRepoForService struct {
	CreateFunc                 func(slot model.Slot) (model.Slot, error)
	GetAvailableByDoctorIDFunc func(doctorID string) ([]model.Slot, error)
	DeleteByIDFunc             func(id string) (int64, error)
	GetByIDFunc                func(id uint) (model.Slot, error)
	UpdateIsBookedFunc         func(id uint, isBooked bool) error
}

func (m *MockSlotRepoForService) Create(slot model.Slot) (model.Slot, error) {
	return m.CreateFunc(slot)
}
func (m *MockSlotRepoForService) GetAvailableByDoctorID(doctorID string) ([]model.Slot, error) {
	return m.GetAvailableByDoctorIDFunc(doctorID)
}
func (m *MockSlotRepoForService) DeleteByID(id string) (int64, error) {
	return m.DeleteByIDFunc(id)
}
func (m *MockSlotRepoForService) GetByID(id uint) (model.Slot, error) {
	return m.GetByIDFunc(id)
}
func (m *MockSlotRepoForService) UpdateIsBooked(id uint, isBooked bool) error {
	return m.UpdateIsBookedFunc(id, isBooked)
}

// Create Slot Test
func TestCreateSlot_Success(t *testing.T) {
	mockRepo := &MockSlotRepoForService{
		CreateFunc: func(slot model.Slot) (model.Slot, error) {
			slot.ID = 1
			return slot, nil
		},
	}

	svc := &SlotService{Repo: mockRepo}
	input := model.Slot{
		StartTime: "2026-05-15 09:00",
		EndTime:   "2026-05-15 10:00",
	}

	created, err := svc.CreateSlot(5, input)

	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if created.DoctorID != 5 {
		t.Errorf("expected DoctorID 5, got %d", created.DoctorID)
	}
}

func TestCreateSlot_MissingStartTime(t *testing.T) {
	svc := &SlotService{Repo: &MockSlotRepoForService{}}
	input := model.Slot{EndTime: "2026-05-15 10:00"}

	_, err := svc.CreateSlot(1, input)

	if err == nil {
		t.Error("expected error when start_time missing, got nil")
	}
	if err.Error() != "ต้องระบุ start_time และ end_time" {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestCreateSlot_MissingEndTime(t *testing.T) {
	svc := &SlotService{Repo: &MockSlotRepoForService{}}
	input := model.Slot{StartTime: "2026-05-15 09:00"}

	_, err := svc.CreateSlot(1, input)
	if err == nil {
		t.Error("expected error when end_time missing, got nil")
	}
}

func TestCreateSlot_InvalidStartTimeFormat(t *testing.T) {
	svc := &SlotService{Repo: &MockSlotRepoForService{}}
	input := model.Slot{
		StartTime: "2026/05/15 09:00",
		EndTime:   "2026-05-15 10:00",
	}

	_, err := svc.CreateSlot(1, input)
	if err == nil {
		t.Error("expected error for invalid start_time format, got nil")
	}
}

func TestCreateSlot_InvalidEndTimeFormat(t *testing.T) {
	svc := &SlotService{Repo: &MockSlotRepoForService{}}
	input := model.Slot{
		StartTime: "2026-05-15 09:00",
		EndTime:   "2026/05/15 10:00",
	}

	_, err := svc.CreateSlot(1, input)
	if err == nil {
		t.Error("expected error for invalid end_time format, got nil")
	}
}

func TestCreateSlot_EndBeforeStart(t *testing.T) {
	svc := &SlotService{Repo: &MockSlotRepoForService{}}
	input := model.Slot{
		StartTime: "2026-05-15 10:00",
		EndTime:   "2026-05-15 09:00",
	}

	_, err := svc.CreateSlot(1, input)
	if err == nil {
		t.Error("expected error when end < start, got nil")
	}
	if err.Error() != "end_time ต้องอยู่หลัง start_time" {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestCreateSlot_EndEqualStart(t *testing.T) {
	svc := &SlotService{Repo: &MockSlotRepoForService{}}
	input := model.Slot{
		StartTime: "2026-05-15 09:00",
		EndTime:   "2026-05-15 09:00",
	}

	_, err := svc.CreateSlot(1, input)
	if err == nil {
		t.Error("expected error when end == start, got nil")
	}
}

func TestCreateSlot_RepoError(t *testing.T) {
	mockRepo := &MockSlotRepoForService{
		CreateFunc: func(slot model.Slot) (model.Slot, error) {
			return slot, errors.New("db error")
		},
	}

	svc := &SlotService{Repo: mockRepo}
	input := model.Slot{
		StartTime: "2026-05-15 09:00",
		EndTime:   "2026-05-15 10:00",
	}

	_, err := svc.CreateSlot(1, input)
	if err == nil {
		t.Error("expected error from repo, got nil")
	}
}

// GetAvailableSlots Test
func TestGetAvailableSlots_Success(t *testing.T) {
	mockRepo := &MockSlotRepoForService{
		GetAvailableByDoctorIDFunc: func(doctorID string) ([]model.Slot, error) {
			return []model.Slot{
				{ID: 1, DoctorID: 1, StartTime: "2026-05-15 09:00", EndTime: "2026-05-15 10:00"},
				{ID: 2, DoctorID: 1, StartTime: "2026-05-15 10:00", EndTime: "2026-05-15 11:00"},
			}, nil
		},
	}

	svc := &SlotService{Repo: mockRepo}
	slots, err := svc.GetAvailableSlots("1")

	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if len(slots) != 2 {
		t.Errorf("expected 2 slots, got %d", len(slots))
	}
}

func TestGetAvailableSlots_RepoError(t *testing.T) {
	mockRepo := &MockSlotRepoForService{
		GetAvailableByDoctorIDFunc: func(doctorID string) ([]model.Slot, error) {
			return nil, errors.New("db error")
		},
	}

	svc := &SlotService{Repo: mockRepo}
	_, err := svc.GetAvailableSlots("1")

	if err == nil {
		t.Error("expected error, got nil")
	}
}

// DeleteSlot Test

func TestDeleteSlot_Success(t *testing.T) {
	mockRepo := &MockSlotRepoForService{
		DeleteByIDFunc: func(id string) (int64, error) {
			return 1, nil
		},
	}

	svc := &SlotService{Repo: mockRepo}
	rows, err := svc.DeleteSlot("1")

	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if rows != 1 {
		t.Errorf("expected 1 row, got %d", rows)
	}
}

func TestDeleteSlot_NotFound(t *testing.T) {
	mockRepo := &MockSlotRepoForService{
		DeleteByIDFunc: func(id string) (int64, error) {
			return 0, nil
		},
	}

	svc := &SlotService{Repo: mockRepo}
	rows, _ := svc.DeleteSlot("9999")

	if rows != 0 {
		t.Errorf("expected 0 rows, got %d", rows)
	}
}

func TestDeleteSlot_RepoError(t *testing.T) {
	mockRepo := &MockSlotRepoForService{
		DeleteByIDFunc: func(id string) (int64, error) {
			return 0, errors.New("db error")
		},
	}

	svc := &SlotService{Repo: mockRepo}
	_, err := svc.DeleteSlot("1")

	if err == nil {
		t.Error("expected error, got nil")
	}
}
