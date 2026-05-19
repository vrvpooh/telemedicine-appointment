package service

import (
	"telemedicine-api/model"
)


type FeedbackRepositoryInterface interface {
	CreateFeedback(f model.Feedback) error
	VerifyDoctor(id string, status bool) error
}

type FeedbackService struct {

	Repo FeedbackRepositoryInterface
}

func (s *FeedbackService) SubmitFeedback(f model.Feedback) error {

	return s.Repo.CreateFeedback(f)
}

func (s *FeedbackService) VerifyDoctorStatus(id string, status bool) error {
	return s.Repo.VerifyDoctor(id, status)
}