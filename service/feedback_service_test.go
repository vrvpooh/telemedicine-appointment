package service

import (
	"errors"
	"telemedicine-api/model"
	"testing"
)

// Mock สำหรับ Feedback และ Verify
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

	// ทดสอบ Logic การเรียกใช้
	err := mockRepo.CreateFeedback(model.Feedback{Rating: 5})
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

	err := mockRepo.VerifyDoctor("999", true)
	if err == nil {
		t.Error("expected error, got nil")
	}
}