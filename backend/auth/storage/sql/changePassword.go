package sqlstore

import (
	"context"
	"fmt"
)

// ChangePassword changes a user's password
func (dbConn *SQLStore) ChangePassword(ctx context.Context, email string, newPassword string) error {

	queryString := fmt.Sprintf(
		`UPDATE %s SET password= $1 WHERE email = $2;`,
		dbConn.registryTN)

	_, err := dbConn.db.ExecContext(ctx, queryString, newPassword, email)

	return err
}
