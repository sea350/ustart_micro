package sqlstore

import (
	"context"
)

// IDLookup looks up if a row exists using a certain uuid
func (dbConn *SQLStore) IDLookup(ctx context.Context, uuid string) (bool, error) {
	rows, err := dbConn.db.QueryContext(ctx, `SELECT "uuid" FROM `+dbConn.RegistryTN+` WHERE uuid= '`+uuid+"';")
	if err != nil {
		return false, err
	}

	defer rows.Close()

	if rows.Next() {
		if err := rows.Scan(&uuid); err != nil {
			return false, err
		}
		if rows.Next() {
			return true, ErrTooManyResults
		}
		return true, nil
	}
	return false, ErrUserDoesNotExist
}
