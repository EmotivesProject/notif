package db

import (
	"context"

	commonMongo "github.com/TomBowyerResearchProject/common/mongo"
	commonNotification "github.com/TomBowyerResearchProject/common/notification"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateNotification(notif *commonNotification.Notification) error {
	_, err := insetIntoCollection(NotificationsCollection, notif)

	return err
}

func insetIntoCollection(collectionName string, document interface{}) (*mongo.InsertOneResult, error) {
	db := commonMongo.GetDatabase()
	collection := db.Collection(collectionName)

	return collection.InsertOne(context.TODO(), document)
}
