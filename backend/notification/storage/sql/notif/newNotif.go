package sqlstore

import (
	"context"
	"fmt"
)

//NewNotif stores a new notification
func (dbConn *SQLStore) NewNotif(ctx context.Context, uuid, activityID string) error {

	queryString := fmt.Sprintf(
		`INSERT INTO %s (uuid, activity_id, seen) 
		VALUES ($1, $2, $3);`,
		dbConn.notifTableName)

	_, err := dbConn.db.ExecContext(ctx, queryString, uuid, activityID, false)

	return err
}
