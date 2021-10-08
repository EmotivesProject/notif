package db

import (
	"context"

	commonNotification "github.com/TomBowyerResearchProject/common/notification"
	commonPostgres "github.com/TomBowyerResearchProject/common/postgres"
)

func CreateNotification(ctx context.Context, notif *commonNotification.Notification) error {
	connection := commonPostgres.GetDatabase()

	return connection.QueryRow(
		ctx,
		`INSERT INTO notifications(username,type,title,message,link,post_id,username_to,created_at,seen)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9) RETURNING id`,
		notif.Username,
		notif.Type,
		notif.Title,
		notif.Message,
		notif.Link,
		notif.PostID,
		notif.UsernameTo,
		notif.CreatedAt,
		notif.Seen,
	).Scan(
		&notif.ID,
	)
}
