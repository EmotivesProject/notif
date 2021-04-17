package db

import (
	"notif/model"
)

func CreateNotification(notif *model.Notification) error {
	connection := GetDB()
	createdNotification := connection.Create(notif)

	if createdNotification.Error != nil {
		return createdNotification.Error
	}
	return nil
}
