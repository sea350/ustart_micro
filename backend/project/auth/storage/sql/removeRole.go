package sqlstore

import (
	"context"
	"fmt"
)

//RemoveRole removes a role from a project
func (dbConn *SQLStore) RemoveRole(ctx context.Context, projectID, roleName string) error {
	queryString := fmt.Sprintf(
		`DELETE FROM %s
		WHERE project_id = $1 AND project_id = $2;`,
		dbConn.memberTN)

	_, err := dbConn.db.ExecContext(ctx, queryString, projectID, roleName)

	return err
}
