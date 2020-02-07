package sqlstore

import (
	"context"
	"fmt"

	"github.com/sea350/ustart_micro/backend/project/projectpb"
)

// GetRoleProfile retreivs a certain role profile
func (dbConn *SQLStore) GetRoleProfile(ctx context.Context, projectID, roleName string) (projectpb.Role, error) {

	queryString := fmt.Sprintf(
		`SELECT manage_members, change_icon, change_banner, manage_entries, manage_links, manage_tags
		FROM %s
		WHERE project_id = $1 AND role_name = $2;`,
		dbConn.roleTN,
	)
	rows, err := dbConn.db.QueryContext(ctx, queryString, projectID, roleName)
	if err != nil {
		return projectpb.Role{}, err
	}

	defer rows.Close()
	var role projectpb.Role

	if rows.Next() { //if the first row loads
		//scan/load data into struct
		if err := rows.Scan(
			&roleName,
			&role.ManageMembers,
			&role.ChangeIcon,
			&role.ChangeBanner,
			&role.ManageEntries,
			&role.ManageLinks,
			&role.ManageTags); err != nil { //unless err!=nil
			return projectpb.Role{}, err
		}

		//if there is a second row there is a problem
		if rows.Next() {
			return projectpb.Role{}, errTooManyResults
		}

	} else {
		//if first row doesnt load there are no results
		return projectpb.Role{}, errNoResultsFound
	}
	return role, nil //everything went well
}
