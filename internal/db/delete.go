package db

import (
	"context"

	commonPostgres "github.com/TomBowyerResearchProject/common/postgres"
)

func DeleteNotificationByPostID(ctx context.Context, postID int) {
	connection := commonPostgres.GetDatabase()
	_, _ = connection.Exec(
		ctx,
		`DELETE FROM notifications WHERE post_id = $1`,
		postID,
	)
}
