package sqlstore

import (
	"context"
	"fmt"
)

//SetToken sets the token and expiration for the acc of the given email
func (dbConn *SQLStore) SetToken(ctx context.Context, email, token, expiration string) error {

	queryString := fmt.Sprintf(
		`UPDATE %s SET token= $1, expiration_date = $2 WHERE email = $3;`,
		dbConn.registryTN)

	_, err := dbConn.db.QueryContext(ctx, queryString, token, expiration, email)

	return err

}
