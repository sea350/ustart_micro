package sqlstore

import (
	"context"
	"fmt"
)

//SetActive sets an existing record as active and updates the login time
func (dbConn *SQLStore) SetActive(ctx context.Context, sessionID string, loginTime string, expiration string) error {
	queryString := fmt.Sprintf(
		`UPDATE %s SET active = $1, active_since = $2, expiration = $3  WHERE session_id = $4;`,
		dbConn.SessionTableName)

	_, err := dbConn.db.QueryContext(ctx, queryString, true, loginTime, expiration, sessionID)

	return err
}
