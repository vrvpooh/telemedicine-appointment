package service

import (
	"telemedicine-api/model"
)


type NotificationRepositoryInterface interface {
	GetNotifications(userID int) ([]model.Notification, error)
}

type NotificationService struct {

	Repo NotificationRepositoryInterface
}

func (s *NotificationService) GetUserNotifications(userID int) ([]model.Notification, error) {
	return s.Repo.GetNotifications(userID)
}