package sqlstore

import (
	"context"
	"fmt"
	"time"
)

//ChangeSeen changes the seen value while also updating timestamp records
func (dbConn *SQLStore) ChangeSeen(ctx context.Context, notifID string, value bool, timestamp time.Time) error {

	queryString := fmt.Sprintf(
		`UPDATE %s
		SET seen = $2, time_seen = $3
		WHERE notif_id = $1;`,
		dbConn.notifTableName)

	_, err := dbConn.db.ExecContext(ctx, queryString, notifID, value, timestamp)

	return err
}
