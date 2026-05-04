package model

// notification_models.go
type Notification struct {
    ID        int    `json:"id"`
    UserID    int    `json:"user_id"`
    Message   string `json:"message"`
    IsRead    int    `json:"is_read"` // 0 = false, 1 = true
    CreatedAt string `json:"created_at"`
}

// feedback_models.go
type Feedback struct {
    ID        int     `json:"id"`
    UserID    int     `json:"user_id"`
    DoctorID  int     `json:"doctor_id"`
    Rating    float64 `json:"rating"`
    Comment   string  `json:"comment"`
}