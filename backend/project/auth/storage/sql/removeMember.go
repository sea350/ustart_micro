package sqlstore

import (
	"context"
	"fmt"
)

//RemoveMember removes a member from a project
func (dbConn *SQLStore) RemoveMember(ctx context.Context, uuid, projectID string) error {
	queryString := fmt.Sprintf(
		`DELETE FROM %s
		WHERE uuid = $1 AND project_id = $2;`,
		dbConn.memberTN)

	_, err := dbConn.db.ExecContext(ctx, queryString, uuid, projectID)

	return err
}
