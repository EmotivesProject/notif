package db

import (
	"context"

	"github.com/TomBowyerResearchProject/common/logger"
	commonMongo "github.com/TomBowyerResearchProject/common/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

func DeleteNotificationByPostID(ctx context.Context, postID int) {
	db := commonMongo.GetDatabase()
	notifCollection := db.Collection(NotificationsCollection)

	_, err := notifCollection.DeleteMany(
		ctx,
		bson.M{
			"post_id": postID,
		},
	)
	if err != nil {
		logger.Error(err)
	}
}
