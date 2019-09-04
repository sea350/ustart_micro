package sqlstore

import (
	"context"
	"fmt"
)

// ChangePassword changes a user's password
func (dbConn *SQLStore) ChangePermissionProfileCategory context.Context, pid string, uuid string, newPassword string), hasPrivilege bool error {
	// pull data attached to the email

	if hasPrivilege{
	queryString := fmt.Sprintf(
		`UPDATE %s SET = '%s' WHERE email = '%s';`,
		dbConn.ProjectDataTN, newPassword, email)

	_, err := dbConn.db.QueryContext(ctx, queryString)

	return err
	}
}
