package sqlstore

import (
	"context"
	"fmt"
)

// GetPassword retreivs a user's password
func (dbConn *SQLStore) GetPassword(ctx context.Context, email string) (string, error) {

	queryString := fmt.Sprintf(
		"SELECT password FROM %s WHERE email = $1;",
		dbConn.registryTN)

	rows, err := dbConn.db.QueryContext(ctx, queryString, email)
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
			return "", errTooManyResults
		}
		return password, nil //everything went well
	}
	return "", errUserDoesNotExist
}
