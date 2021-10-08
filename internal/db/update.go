package db

import (
	"context"

	commonPostgres "github.com/TomBowyerResearchProject/common/postgres"
)

func UpdateNotificationsSeen(ctx context.Context, link, username string) {
	connection := commonPostgres.GetDatabase()

	_, _ = connection.Exec(
		ctx,
		"UPDATE notifications SET seen = true WHERE username = $1 AND link = $2",
		username,
		link,
	)
}

func UpdateNotificationID(ctx context.Context, id int) {
	connection := commonPostgres.GetDatabase()

	_, _ = connection.Exec(
		ctx,
		"UPDATE notifications SET seen = true WHERE id = $1",
		id,
	)
}
