package sqlstore

import (
	"context"
	"fmt"
)

//FindSession finds a session using uuid and ip adress, returns sessionID
func (dbConn *SQLStore) FindSession(ctx context.Context, uuid string, ipAddress string) (string, error) {
	rows, err := dbConn.db.QueryContext(ctx, fmt.Sprintf("SELECT session_id FROM %s WHERE uuid = $1 AND ipAddress = $2;",
		dbConn.SessionTableName), uuid, ipAddress)
	if err != nil {
		return "", err
	}

	defer rows.Close()

	var sessionID string
	if rows.Next() { //calls the first row
		if err := rows.Scan(&sessionID); err != nil {
			return "", err
		}

		//IGNORE EXTRANEOUS RESULTS, THESE DONT MATTER

		// if rows.Next() { //if there is a second row after the first an error has occured
		// 	return "", errTooManyResults
		// }

		return sessionID, nil //everything went well
	}
	return sessionID, errSessionDoesNotExist
}
