package db

import (
	"context"
	"notif/model"

	"github.com/TomBowyerResearchProject/common/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	paged = 5
)

func FindNotificationsByUsername(username string, pageOffset int64) *[]model.Notification {
	var notifications []model.Notification

	query := bson.M{"username": username}

	findOptions := options.Find()
	findOptions.SetSort(bson.M{"created_at": -1})
	findOptions.SetSkip(pageOffset * paged)
	findOptions.SetLimit(paged)

	db := GetDatabase()
	notifCollection := db.Collection(NotificationsCollection)
	cursor, err := notifCollection.Find(context.TODO(), query, findOptions)
	if err == mongo.ErrNoDocuments {
		return &notifications
	}

	for cursor.Next(context.TODO()) {
		//Create a value into which the single document can be decoded
		var notification model.Notification
		err := cursor.Decode(&notification)
		if err != nil {
			logger.Error(err)
			continue
		}

		notifications = append(notifications, notification)
	}

	return &notifications
}
