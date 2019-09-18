package sqlstore

import (
	"context"
	"fmt"
)

//ModifyMemberRole changes member role from a project
func (dbConn *SQLStore) ModifyMemberRole(ctx context.Context, uuid, projectID, newRole string) error {
	queryString := fmt.Sprintf(
		`UPDATE %s
		SET role_name = $3
		WHERE uuid = $1 AND project_id = $2;`,
		dbConn.memberTN)

	_, err := dbConn.db.ExecContext(ctx, queryString, uuid, projectID, newRole)

	return err
}
