package service

import (
	"telemedicine-api/model"
	"telemedicine-api/repository"
)

type NotificationService struct {
    Repo *repository.NotificationRepository
}

func (s *NotificationService) GetUserNotifications(userID int) ([]model.Notification, error) {
    return s.Repo.GetNotifications(userID)
}