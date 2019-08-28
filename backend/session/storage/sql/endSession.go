package sqlstore

import (
	"context"
	"fmt"
)

//EndSession marks an existing session inactive
func (dbConn *SQLStore) EndSession(ctx context.Context, sessionID string, expiration string) error {
	queryString := fmt.Sprintf(
		`UPDATE %s SET active = $1, expiration = $2 WHERE session_id = $3;`,
		dbConn.SessionTableName)

	_, err := dbConn.db.QueryContext(ctx, queryString, false, expiration, sessionID)

	return err
}
