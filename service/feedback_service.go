package service

import (
	"telemedicine-api/model"
	"telemedicine-api/repository"
)

type FeedbackService struct {
    Repo *repository.FeedbackRepository
}

func (s *FeedbackService) SubmitFeedback(f model.Feedback) error {
    // สามารถเพิ่ม logic ตรวจสอบ rating 1-5 ที่นี่ได้
    return s.Repo.CreateFeedback(f)
}

func (s *FeedbackService) VerifyDoctorStatus(id string, status bool) error {
    return s.Repo.VerifyDoctor(id, status)
}