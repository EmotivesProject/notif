package db

import (
	"context"

	"github.com/TomBowyerResearchProject/common/logger"
	commonMongo "github.com/TomBowyerResearchProject/common/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdateNotificationsSeen(ctx context.Context, link, username string) {
	db := commonMongo.GetDatabase()
	notifCollection := db.Collection(NotificationsCollection)

	_, err := notifCollection.UpdateMany(
		ctx,
		bson.M{
			"link":     link,
			"username": username,
		},
		bson.D{
			{"$set", bson.D{{"seen", true}}}, //nolint:govet
		},
	)
	if err != nil {
		logger.Error(err)
	}
}

func UpdateNotificationID(ctx context.Context, id primitive.ObjectID) {
	db := commonMongo.GetDatabase()
	notifCollection := db.Collection(NotificationsCollection)

	_, err := notifCollection.UpdateMany(
		ctx,
		bson.M{
			"_id": id,
		},
		bson.D{
			{"$set", bson.D{{"seen", true}}}, //nolint:govet
		},
	)
	if err != nil {
		logger.Error(err)
	}
}
