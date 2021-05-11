package db

import (
	"context"

	commonMongo "github.com/TomBowyerResearchProject/common/mongo"
	commonNotification "github.com/TomBowyerResearchProject/common/notification"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateNotification(ctx context.Context, notif *commonNotification.Notification) error {
	_, err := insetIntoCollection(ctx, NotificationsCollection, notif)

	return err
}

func insetIntoCollection(
	ctx context.Context,
	collectionName string,
	document interface{},
) (*mongo.InsertOneResult, error) {
	db := commonMongo.GetDatabase()
	collection := db.Collection(collectionName)

	return collection.InsertOne(ctx, document)
}
