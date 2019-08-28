package sqlstore

import (
	"context"
	"fmt"
)

// ChangePassword changes a user's password
func (dbConn *SQLStore) ChangePassword(ctx context.Context, email string, newPassword string) error {

	queryString := fmt.Sprintf(
		`UPDATE %s SET password= '%s' WHERE email = '%s';`,
		dbConn.RegistryTN, newPassword, email)

	_, err := dbConn.db.QueryContext(ctx, queryString)

	return err
}
