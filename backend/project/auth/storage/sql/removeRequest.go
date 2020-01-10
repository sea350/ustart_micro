package sqlstore

import (
	"context"
	"fmt"
)

//RemoveRequest removes a request from a project
func (dbConn *SQLStore) RemoveRequest(ctx context.Context, uuid, projectID string) error {
	queryString := fmt.Sprintf(
		`DELETE FROM %s
		WHERE uuid = $1 AND project_id = $2;`,
		dbConn.requestTN)

	_, err := dbConn.db.ExecContext(ctx, queryString, uuid, projectID)

	return err
}
