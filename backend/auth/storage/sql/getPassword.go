package sqlstore

import (
	"context"
)

// GetPassword retreivs a user's password
func (dbConn *SQLStore) GetPassword(ctx context.Context, email string) (string, error) {

	rows, err := dbConn.db.QueryContext(ctx, "SELECT password FROM "+dbConn.RegistryTN+" WHERE email = '"+email+"';")
	if err != nil {
		return "", err
	}

	defer rows.Close()

	var password string
	if rows.Next() { //calls the first row
		if err := rows.Scan(&password); err != nil {
			return "", err
		}
		if rows.Next() { //if there is a second row after the first an error has occured
			return "", ErrTooManyResults
		}
		return password, nil //everything went well
	}
	return "", ErrUserDoesNotExist
}
