package sqlstore

import (
	"context"
	"database/sql"
)

// GetPassword retreivs a user's password
func (dbConn *SQLStore) GetPassword(ctx context.Context, email string) (string, error) {

	rows, err := dbConn.db.QueryContext(ctx, "SELECT password FROM "+dbConn.RegistryTN+" WHERE email = "+email+";")
	defer rows.Close()

	// if there are no hits, then no one exists by that email
	if err == sql.ErrNoRows {
		return "", ErrUserDoesNotExist
	}
	if err != nil {
		return "", err
	}

	var password string
	for rows.Next() {
		// if it moves on to the next row and a password has already been found,
		// then something went wrong in registry
		if password != "" {
			return "", ErrTooManyResults
		}
		if err := rows.Scan(&password); err != nil {
			return "", err
		}
	}

	return password, nil
}
