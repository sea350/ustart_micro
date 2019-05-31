package sqlstore

import (
	"context"
)

// Lookup looks up if a document exists using a certain email. It returns a uuid and an error
func (dbConn *SQLStore) Lookup(ctx context.Context, email string) (string, error) {
	rows, err := dbConn.db.QueryContext(ctx, `SELECT "uuid" FROM `+dbConn.RegistryTN+` WHERE email= '`+email+"';")
	if err != nil {
		return "", err
	}

	defer rows.Close()

	var uuid string
	if rows.Next() {
		if err := rows.Scan(&uuid); err != nil {
			return "", err
		}
		if rows.Next() {
			return uuid, errTooManyResults
		}
		return uuid, nil
	}
	return "", errUserDoesNotExist
}
