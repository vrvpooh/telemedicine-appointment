package repository

import (
	"telemedicine-api/config"
	"telemedicine-api/model"
)

type FeedbackRepository struct{}

func (r *FeedbackRepository) CreateFeedback(f model.Feedback) error {
    _, err := config.DB.Exec("INSERT INTO feedbacks (user_id, doctor_id, rating, comment) VALUES (?, ?, ?, ?)",
        f.UserID, f.DoctorID, f.Rating, f.Comment)
    return err
}

func (r *FeedbackRepository) VerifyDoctor(id string, status bool) error {
    val := 0
    if status { val = 1 }
    _, err := config.DB.Exec("UPDATE doctors SET is_verified = ? WHERE id = ?", val, id)
    return err
}