package sqlstore

import (
	"context"
)

// ChangePassword changes a user's password
func (dbConn *SQLStore) ChangePassword(ctx context.Context, email string, newPassword string) error {
	// pull data attached to the email
	_, err := dbConn.db.QueryContext(ctx, "UPDATE "+dbConn.RegistryTN+" SET password = "+newPassword+" WHERE email = "+email+";")

	return err
}
