package sqlstore

import (
	"context"
	"fmt"

	"github.com/sea350/ustart_micro/backend/project/auth/types"
)

// GetRoleProfile retreivs a certain role profile
func (dbConn *SQLStore) GetRoleProfile(ctx context.Context, projectID, roleName string) (types.Role, error) {

	queryString := fmt.Sprintf(
		`SELECT role_name, manage_members, change_icon, change_banner, manage_entries, manage_links, manage_tags
		FROM %s
		WHERE project_id = $1;`,
		dbConn.roleTN,
	)
	rows, err := dbConn.db.QueryContext(ctx, queryString, projectID)
	if err != nil {
		return types.Role{}, err
	}

	defer rows.Close()
	var role types.Role

	if rows.Next() { //loop through rows

		if err := rows.Scan(
			&role.Name,
			&role.ManageMembers,
			&role.ChangeIcon,
			&role.ChangeBanner,
			&role.ManageEntries,
			&role.ManageLinks,
			&role.ManageTags); err != nil {
			return types.Role{}, err
		}

	} else {
		return types.Role{}, errNoResultsFound
	}
	return role, nil //everything went well
}
