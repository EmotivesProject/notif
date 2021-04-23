package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Notification struct {
	ID        primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Username  string             `bson:"username" json:"username"`
	Title     string             `bson:"title" json:"title"`
	Message   string             `bson:"message" json:"message"`
	Link      string             `bson:"link" json:"link"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	Seen      bool               `bson:"seen" json:"seen"`
}

// Might need to make this smarter, not sure
func CreateNotificationFromMap(jsonMap map[string]interface{}) Notification {
	return Notification{
		ID:        primitive.NewObjectID(),
		Username:  jsonMap["username"].(string),
		Title:     jsonMap["title"].(string),
		Message:   jsonMap["message"].(string),
		Link:      jsonMap["link"].(string),
		CreatedAt: time.Now(),
		Seen:      false,
	}
}
