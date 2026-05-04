package repository

import (
	"telemedicine-api/config"
	"telemedicine-api/model"
)

type NotificationRepository struct{}

func (r *NotificationRepository) GetNotifications(userID int) ([]model.Notification, error) {
    rows, err := config.DB.Query("SELECT id, user_id, message, is_read, created_at FROM notifications WHERE user_id = ?", userID)
    if err != nil { return nil, err }
    defer rows.Close()

    var notifications []model.Notification
    for rows.Next() {
        var n model.Notification
        rows.Scan(&n.ID, &n.UserID, &n.Message, &n.IsRead, &n.CreatedAt)
        notifications = append(notifications, n)
    }
    return notifications, nil
}