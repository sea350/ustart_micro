package sqlstore

import (
	"context"
	"fmt"
)

//AcceptRequest removes the request from the request list, but adds the new member to the project
func (dbConn *SQLStore) AcceptRequest(ctx context.Context, uuid, projectID, requestDate string) error {
	queryString := fmt.Sprintf(
		`INSERT INTO %s (uuid, project_id, join_date) 
		VALUES ($1, $2, $3);`,
		dbConn.memberTN)

	_, err := dbConn.db.ExecContext(ctx, queryString, uuid, projectID, joinDate)

	return err
}
