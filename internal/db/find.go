package db

import (
	"notif/model"
)

func FindNotificationsByUsername(username string, offset int64) []model.Notification {
	var notifications []model.Notification

	database.Where("username = ? ORDER BY created_at LIMIT 5 OFFSET ?", username, offset).Find(&notifications)

	return notifications
}
