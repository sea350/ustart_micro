package sqlstore

import (
	"context"
	"fmt"
	"time"
)

//NewActivity enters a new log entry
func (dbConn *SQLStore) NewActivity(ctx context.Context, uuid, objectID, action string, timestamp time.Time) (string, error) {

	queryString := fmt.Sprintf(
		`INSERT INTO %s (actor_id, object_id, action, time) 
		VALUES ($1, $2, $3, $4)
		RETURNING activity_id;`,
		dbConn.activityTableName)

	var activityID string
	err := dbConn.db.QueryRowContext(ctx, queryString, uuid, objectID, action, timestamp).Scan(&activityID)

	return activityID, err
}
