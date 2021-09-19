package db

import (
	"context"
	"errors"

	"github.com/TomBowyerResearchProject/common/logger"
	commonMongo "github.com/TomBowyerResearchProject/common/mongo"
	commonNotification "github.com/TomBowyerResearchProject/common/notification"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	paged = 5
)

func FindNotificationsByUsername(
	ctx context.Context,
	username string,
	pageOffset int64,
) []commonNotification.Notification {
	notifications := make([]commonNotification.Notification, 0)

	query := bson.M{"username": username}

	findOptions := options.Find()
	findOptions.SetSort(bson.M{"created_at": -1})
	findOptions.SetSkip(pageOffset * paged)
	findOptions.SetLimit(paged)

	db := commonMongo.GetDatabase()
	notifCollection := db.Collection(NotificationsCollection)

	cursor, err := notifCollection.Find(ctx, query, findOptions)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return notifications
	}

	for cursor.Next(ctx) {
		// Create a value into which the single document can be decoded.
		var notification commonNotification.Notification

		err := cursor.Decode(&notification)
		if err != nil {
			logger.Error(err)

			continue
		}

		notifications = append(notifications, notification)
	}

	return notifications
}
