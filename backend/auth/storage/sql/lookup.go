package sqlstore

import (
	"context"
)

// Lookup looks up if a document exists using a certain email
func (dbConn *SQLStore) Lookup(ctx context.Context, email string) (bool, error) {
	rows, err := dbConn.db.QueryContext(ctx, `SELECT "uuid" FROM `+dbConn.RegistryTN+` WHERE email= '`+email+"';")
	if err != nil {
		return false, err
	}

	defer rows.Close()

	var uuid string
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
