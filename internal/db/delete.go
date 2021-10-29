package db

import (
	"context"

	commonPostgres "github.com/EmotivesProject/common/postgres"
)

func DeleteNotificationByPostID(ctx context.Context, postID int) {
	connection := commonPostgres.GetDatabase()
	_, _ = connection.Exec(
		ctx,
		`DELETE FROM notifications WHERE post_id = $1`,
		postID,
	)
}

func DeleteNotificationByPostIDUsernameAndType(ctx context.Context, postID int, username, notifType string) {
	connection := commonPostgres.GetDatabase()
	_, _ = connection.Exec(
		ctx,
		`DELETE FROM notifications WHERE post_id = $1 AND username_to = $2 AND type = $3`,
		postID,
		username,
		notifType,
	)
}
