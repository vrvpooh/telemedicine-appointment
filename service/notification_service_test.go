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
		Repo: mockRepo,
	}


	data, err := svc.GetUserNotifications(1)
	
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if len(data) != 1 {
		t.Errorf("expected 1 item, got %d", len(data))
	}
}

func TestGetUserNotifications_Error(t *testing.T) {
	mockRepo := &MockNotificationRepo{
		GetNotificationsFunc: func(userID int) ([]model.Notification, error) {
			return nil, errors.New("database error")
		},
	}


	svc := &NotificationService{
		Repo: mockRepo,
	}


	_, err := svc.GetUserNotifications(1)
	if err == nil {
		t.Error("expected error, got nil")
	}
}