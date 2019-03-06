package sqlstore

import (
	"context"
	"database/sql"
)

// Lookup looks up if a document exists using a certain email
func (dbConn *SQLStore) Lookup(ctx context.Context, email string) (bool, error) {
	rows, err := dbConn.db.QueryContext(ctx, "SELECT uuid FROM "+dbConn.RegistryTN+" WHERE email="+email+";")
	defer rows.Close()

	// if there are no hits, then no one exists by that email
	if err == sql.ErrNoRows {
		return false, ErrUserDoesNotExist
	}
	if err != nil {
		return false, err
	}

	var uuid string
	for rows.Next() {
		// if it moves on to the next row and a uuid has already been found,
		// then something went wrong in registry
		if uuid != "" {
			return true, ErrTooManyResults
		}
		if err := rows.Scan(&uuid); err != nil {
			return false, err
		}
	}
	return true, nil
}
