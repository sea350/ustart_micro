package sqlstore

import (
	"context"
	"fmt"
)

//NewSession saves a new session record, returns the session id of new record
func (dbConn *SQLStore) NewSession(ctx context.Context, uuid string, ipAddress string, username string, loginTime string, expiration string) (string, error) {
	queryString := fmt.Sprintf(
		`INSERT INTO %s (uuid, ip_address, active_since, expiration, active) 
		VALUES ( $1, $2, $3, $4, $5)
		RETURNING session_id;`,
		dbConn.SessionTableName)

	var sessionID string
	err := dbConn.db.QueryRowContext(ctx, queryString, uuid, ipAddress, loginTime, expiration, true).Scan(&sessionID)

	return sessionID, err
}
