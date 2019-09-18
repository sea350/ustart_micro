package sqlstore

import (
	"context"
	"fmt"
)

// GetMemberRole retrieves the role from a certain member
func (dbConn *SQLStore) GetMemberRole(ctx context.Context, projectID, memberID string) (string, error) {

	queryString := fmt.Sprintf(
		`SELECT role_name
		FROM %s
		WHERE project_id = $1;`,
		dbConn.memberTN,
	)
	rows, err := dbConn.db.QueryContext(ctx, queryString, projectID)
	if err != nil {
		return "", err
	}

	defer rows.Close()
	var roleName string

	if rows.Next() { //loop through rows
		if err := rows.Scan(&roleName); err != nil {
			return "", err
		}
		if rows.Next() {
			return "", errTooManyResults
		}

	} else {
		return "", err
	}

	return roleName, nil //everything went well
}
