package db

import (
	"context"

	commonNotification "github.com/TomBowyerResearchProject/common/notification"
	commonPostgres "github.com/TomBowyerResearchProject/common/postgres"
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

	connection := commonPostgres.GetDatabase()

	actualOffset := pageOffset * paged

	rows, err := connection.Query(
		ctx,
		`SELECT * FROM notifications
		WHERE username = $1
		ORDER BY created_at desc LIMIT $3 OFFSET $2`,
		username,
		actualOffset,
		paged,
	)
	if err != nil {
		return notifications
	}

	for rows.Next() {
		var notif commonNotification.Notification

		err := rows.Scan(
			&notif.ID,
			&notif.Username,
			&notif.Type,
			&notif.Title,
			&notif.Message,
			&notif.Link,
			&notif.PostID,
			&notif.UsernameTo,
			&notif.CreatedAt,
			&notif.Seen,
		)
		if err != nil {
			continue
		}

		notifications = append(notifications, notif)
	}

	return notifications
}
