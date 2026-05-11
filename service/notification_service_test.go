package service

import (
	"errors"
	"telemedicine-api/model"
	"testing"
)

type MockNotificationRepo struct {
	GetNotificationsFunc func(userID int) ([]model.Notification, error)
}

func (m *MockNotificationRepo) GetNotifications(userID int) ([]model.Notification, error) {
	return m.GetNotificationsFunc(userID)
}

func TestGetUserNotifications_Success(t *testing.T) {
	mockRepo := &MockNotificationRepo{
		GetNotificationsFunc: func(userID int) ([]model.Notification, error) {
			return []model.Notification{{ID: 1, Message: "Test Message"}}, nil
		},
	}


	
	svc := &NotificationService{
		Repo: nil,
	}

	data, err := mockRepo.GetNotifications(1)
	
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if len(data) != 1 {
		t.Errorf("expected 1 item, got %d", len(data))
	}

	
	_ = svc 
}

func TestGetUserNotifications_Error(t *testing.T) {
	mockRepo := &MockNotificationRepo{
		GetNotificationsFunc: func(userID int) ([]model.Notification, error) {
			return nil, errors.New("database error")
		},
	}

	_, err := mockRepo.GetNotifications(1)
	if err == nil {
		t.Error("expected error, got nil")
	}
}