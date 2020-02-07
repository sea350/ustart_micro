package sqlstore

import (
	"context"
	"fmt"

	"github.com/sea350/ustart_micro/backend/project/projectpb"
)

// GetMembers retreivs all members from a certain project
func (dbConn *SQLStore) GetMembers(ctx context.Context, projectID string) ([]projectpb.Member, error) {

	queryString := fmt.Sprintf(
		`SELECT uuid, role_name
		FROM %s
		WHERE project_id = $1;`,
		dbConn.memberTN,
	)
	rows, err := dbConn.db.QueryContext(ctx, queryString, projectID)
	if err != nil {
		return []projectpb.Member{}, err
	}

	defer rows.Close()
	var mems []projectpb.Member

	for rows.Next() { //loop through rows
		var uuid, roleName string
		if err := rows.Scan(&uuid, &roleName); err != nil {
			return []projectpb.Member{}, err
		}
		mems = append(mems, projectpb.Member{
			UUID:     uuid,
			RoleName: roleName,
		})
	}
	if len(mems) == 0 {
		return []projectpb.Member{}, errNoResultsFound
	}

	return mems, nil //everything went well
}
