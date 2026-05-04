package model

import "time"

type Appointment struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	PatientID uint      `json:"patient_id" binding:"required"`
	SlotID    uint      `json:"slot_id" binding:"required"`
	Status    string    `json:"status" gorm:"default:pending"`
	ZoomToken string    `json:"zoom_token"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
