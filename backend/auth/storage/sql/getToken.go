package sqlstore

import (
	"context"
	"fmt"
)

//GetToken looks up the token and expiration date associated with the given email, returns token and expiration
func (dbConn *SQLStore) GetToken(ctx context.Context, email string) (string, string, error) {
	queryString := fmt.Sprintf(
		`SELECT token, expiration_date FROM %s WHERE email= $1;`,
		dbConn.registryTN)

	rows, err := dbConn.db.QueryContext(ctx, queryString, email)
	if err != nil {
		return "", "", err
	}

	defer rows.Close()

	var token string
	var stringDate string
	if rows.Next() {
		if err := rows.Scan(&token, &stringDate); err != nil {
			return "", "", err
		}
		if rows.Next() {
			return "", "", errTooManyResults
		}
		return token, stringDate, err
	}
	return "", "", errUserDoesNotExist
}
