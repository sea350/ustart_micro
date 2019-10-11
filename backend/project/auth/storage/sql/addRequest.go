package sqlstore

import (
	"context"
	"fmt"
)

//AddRequest stores a new join request given their information
func (dbConn *SQLStore) AddRequest(ctx context.Context, uuid, projectID, requestDate string) error {
	queryString := fmt.Sprintf(
		`INSERT INTO %s (uuid, project_id, join_date) 
		VALUES ($1, $2, $3);`,
		dbConn.memberTN)

	_, err := dbConn.db.ExecContext(ctx, queryString, uuid, projectID, joinDate)

	return err
}
