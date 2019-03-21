package sqlstore

import (
	"context"
	"time"
)

// GetToken looks up the token and expiration date associated with the given email
func (dbConn *SQLStore) GetToken(ctx context.Context, email string) (string, time.Time, error) {
	rows, err := dbConn.db.QueryContext(ctx, `SELECT "token", "expiration_date" FROM `+dbConn.RegistryTN+` WHERE email= '`+email+"';")
	if err != nil {
		return "", time.Time{}, err
	}

	defer rows.Close()

	var token string
	var stringDate string
	if rows.Next() {
		if err := rows.Scan(&token, &stringDate); err != nil {
			return "", time.Time{}, err
		}
		if rows.Next() {
			return "", time.Time{}, ErrTooManyResults
		}
		t, err := time.Parse(dbConn.TimeFormat, stringDate)
		return token, t, err
	}
	return "", time.Time{}, ErrUserDoesNotExist
}
