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

func (n Notification) FillNotification(username string) Notification {
	n.ID = primitive.NewObjectID()
	n.Username = username
	n.Seen = false
	n.CreatedAt = time.Now()
	return n
}
