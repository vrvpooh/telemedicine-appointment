package service

import (
	"errors"
	"telemedicine-api/model"
	"testing"
)


type MockFeedbackRepo struct {
	CreateFeedbackFunc func(f model.Feedback) error
	VerifyDoctorFunc   func(id string, status bool) error
}

func (m *MockFeedbackRepo) CreateFeedback(f model.Feedback) error {
	return m.CreateFeedbackFunc(f)
}
func (m *MockFeedbackRepo) VerifyDoctor(id string, status bool) error {
	return m.VerifyDoctorFunc(id, status)
}

func TestSubmitFeedback_Success(t *testing.T) {
	mockRepo := &MockFeedbackRepo{
		CreateFeedbackFunc: func(f model.Feedback) error {
			return nil
		},
	}

	
	svc := &FeedbackService{Repo: mockRepo}

	err := svc.SubmitFeedback(model.Feedback{Rating: 5})
	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}
}

func TestVerifyDoctorStatus_Error(t *testing.T) {
	mockRepo := &MockFeedbackRepo{
		VerifyDoctorFunc: func(id string, status bool) error {
			return errors.New("doctor not found")
		},
	}

	svc := &FeedbackService{Repo: mockRepo}

	err := svc.VerifyDoctorStatus("999", true)
	if err == nil {
		t.Error("expected error, got nil")
	}
}