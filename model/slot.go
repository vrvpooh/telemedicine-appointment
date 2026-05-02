package model

type Slot struct {
	ID        int64  `json:"id"`
	DoctorID  int64  `json:"doctor_id"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
	IsBooked  bool   `json:"is_booked"`
}
