package sqlstore

import (
	"context"
	"fmt"
)

// IDLookup looks up if a row exists using a certain uuid
func (dbConn *SQLStore) IDLookup(ctx context.Context, uuid string) (bool, error) {
	queryString := fmt.Sprintf(
		`SELECT uuid FROM %s WHERE uuid= $1;`,
		dbConn.registryTN)

	rows, err := dbConn.db.QueryContext(ctx, queryString, uuid)
	if err != nil {
		return false, err
	}

	defer rows.Close()

	if rows.Next() {
		if err := rows.Scan(&uuid); err != nil {
			return false, err
		}
		if rows.Next() {
			return true, errTooManyResults
		}
		return true, nil
	}
	return false, errUserDoesNotExist
}
