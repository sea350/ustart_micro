package sqlstore

import (
	"context"
	"database/sql"
)

// GetPassword retreivs a user's password
func (dbConn *SQLStore) GetPassword(ctx context.Context, email string) (string, error) {

	rows, err := dbConn.db.QueryContext(ctx, "SELECT password FROM "+dbConn.RegistryTN+" WHERE email = '"+email+"';")

	// if there are no hits, then no one exists by that email
	if err == sql.ErrNoRows {
		return "", ErrUserDoesNotExist
	}
	if err != nil {
		return "", err
	}
	defer rows.Close()

	var password string
	if rows.Next() {
		if err := rows.Scan(&password); err != nil {
			return "", err
		}
		if rows.Next() {
			return "", ErrTooManyResults
		} else {
			return password, nil
		}
	}

	return "", ErrUserDoesNotExist
}
