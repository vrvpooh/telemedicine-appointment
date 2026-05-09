package service

import (
	"errors"
	"time"

	"telemedicine-api/model"
	"telemedicine-api/repository"
)

// Interface สำหรับ handler ใช้ mock -> test ได้
type ISlotService interface {
	CreateSlot(doctorID int64, slot model.Slot) (model.Slot, error)
	GetAvailableSlots(doctorID string) ([]model.Slot, error)
	DeleteSlot(id string) (int64, error)
}

type SlotService struct {
	Repo repository.ISlotRepository
}

const slotTimeLayout = "2006-01-02 15:04"

// Create Slot
func (s *SlotService) CreateSlot(doctorID int64, slot model.Slot) (model.Slot, error) {
	slot.DoctorID = doctorID

	if slot.StartTime == "" || slot.EndTime == "" {
		return slot, errors.New("ต้องระบุ start_time และ end_time")
	}

	start, err := time.Parse(slotTimeLayout, slot.StartTime)
	if err != nil {
		return slot, errors.New("รูปแบบ start_time ไม่ถูกต้อง (ต้องเป็น YYYY-MM-DD HH:MM)")
	}
	end, err := time.Parse(slotTimeLayout, slot.EndTime)
	if err != nil {
		return slot, errors.New("รูปแบบ end_time ไม่ถูกต้อง (ต้องเป็น YYYY-MM-DD HH:MM)")
	}

	if !end.After(start) {
		return slot, errors.New("end_time ต้องอยู่หลัง start_time")
	}

	return s.Repo.Create(slot)
}

// GetAvailableSlots
func (s *SlotService) GetAvailableSlots(doctorID string) ([]model.Slot, error) {
	return s.Repo.GetAvailableByDoctorID(doctorID)
}

// DeleteSlot
func (s *SlotService) DeleteSlot(id string) (int64, error) {
	return s.Repo.DeleteByID(id)
}
