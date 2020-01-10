package sqlstore

import (
	"context"
	"fmt"
)

//AddMember stores a new member given their information
func (dbConn *SQLStore) AddMember(ctx context.Context, uuid, projectID, role, joinDate string) error {
	queryString := fmt.Sprintf(
		`INSERT INTO %s (uuid, project_id, role_name, join_date) 
		VALUES ($1, $2, $3, $4);`,
		dbConn.memberTN)

	_, err := dbConn.db.ExecContext(ctx, queryString, uuid, projectID, role, joinDate)

	return err
}
