package db

import (
	"context"
	"notif/model"

	"go.mongodb.org/mongo-driver/mongo"
)

func CreateNotification(notif *model.Notification) error {
	_, err := insetIntoCollection(NotificationsCollection, notif)
	return err
}

func insetIntoCollection(collectionName string, document interface{}) (*mongo.InsertOneResult, error) {
	db := GetDatabase()
	collection := db.Collection(collectionName)
	return collection.InsertOne(context.TODO(), document)
}
