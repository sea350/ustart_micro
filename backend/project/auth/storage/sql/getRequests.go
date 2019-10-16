package sqlstore

import (
	"context"
	"fmt"

	"github.com/sea350/ustart_micro/backend/project/auth/types"
)

// GetRequests retreivs all join requests from a certain project
func (dbConn *SQLStore) GetRequests(ctx context.Context, projectID string) ([]string, error) {

	queryString := fmt.Sprintf(
		`SELECT uuid
		FROM %s
		WHERE project_id = $1;`,
		dbConn.requestTN,
	)

	//Manage members check
	
	rows, err := dbConn.db.QueryContext(ctx, queryString, projectID)
	if err != nil {
		return []string, err
	}

	defer rows.Close()
	var reqs []string

	for rows.Next() { //loop through rows
		var uuid string
		if err := rows.Scan(&uuid); err != nil {
			return []string, err
		}
		reqs = append(reqs, uuid)
	}
	// if len(reqs) == 0 {
	// 	return []string, errNoResultsFound
	// }

	return reqs, nil //everything went well
}
