package sqlstore

import (
	"context"
	"fmt"
)

//Validate sets valid field for the user with the given email
func (dbConn *SQLStore) Validate(ctx context.Context, email string, valid bool) error {

	queryString := fmt.Sprintf(
		`UPDATE %s SET verified= '%t' WHERE email = "%s";`,
		dbConn.RegistryTN, true, email)

	_, err := dbConn.db.QueryContext(ctx, queryString)

	return err

}
