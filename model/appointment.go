package model

type Appointment struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	PatientID uint   `json:"patient_id" binding:"required"`
	SlotID    uint   `json:"slot_id" binding:"required"`
	Status    string `json:"status" gorm:"default:pending"`
	ZoomToken string `json:"zoom_token"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
