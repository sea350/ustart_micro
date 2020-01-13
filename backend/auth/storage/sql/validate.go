package sqlstore

import (
	"context"
	"fmt"
)

//Validate sets valid field for the user with the given email
func (dbConn *SQLStore) Validate(ctx context.Context, email string, valid bool) error {

	queryString := fmt.Sprintf(
		`UPDATE %s SET verified = $1 WHERE email = $2;`,
		dbConn.registryTN)

	_, err := dbConn.db.QueryContext(ctx, queryString, true, email)

	return err

}
