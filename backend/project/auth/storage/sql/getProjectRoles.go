package sqlstore

import (
	"context"
	"fmt"

	"github.com/sea350/ustart_micro/backend/project/projectpb"
)

// GetProjectRoles retreivs all roles from a certain project
func (dbConn *SQLStore) GetProjectRoles(ctx context.Context, projectID string) ([]*projectpb.Role, error) {

	queryString := fmt.Sprintf(
		`SELECT role_name, manage_members, change_icon, change_banner, manage_entries, manage_links, manage_tags
		FROM %s
		WHERE project_id = $1;`,
		dbConn.roleTN,
	)
	rows, err := dbConn.db.QueryContext(ctx, queryString, projectID)
	if err != nil {
		return []*projectpb.Role{}, err
	}

	defer rows.Close()
	var roles []*projectpb.Role

	for rows.Next() { //loop through rows
		var roleName string
		var manageMembers, changeIcon, changeBanner, manageEntries, manageLinks, manageTags bool
		if err := rows.Scan(&roleName, &manageMembers, &changeIcon, &changeBanner, &manageEntries, &manageLinks, &manageTags); err != nil {
			return []*projectpb.Role{}, err
		}
		roles = append(roles, &projectpb.Role{
			Name:          roleName,
			ManageMembers: manageMembers,
			ChangeIcon:    changeIcon,
			ChangeBanner:  changeBanner,
			ManageEntries: manageEntries,
			ManageLinks:   manageLinks,
			ManageTags:    manageTags,
		})
	}
	if len(roles) == 0 {
		return []*projectpb.Role{}, errNoResultsFound
	}

	return roles, nil //everything went well
}
