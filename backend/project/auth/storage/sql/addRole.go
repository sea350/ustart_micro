package sqlstore

import (
	"context"
	"fmt"
)

//AddRole stores a new role for a project
func (dbConn *SQLStore) AddRole(ctx context.Context, role, projectID string, member, icon, banner, entries, links, tags bool) error {
	queryString := fmt.Sprintf(
		`INSERT INTO %s (role_name,
			project_id,
			manage_members,
			change_icon,
			change_banner,
			manage_entries,
			manage_links,
			manage_tags) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8);`,
		dbConn.roleTN)

	_, err := dbConn.db.ExecContext(ctx, queryString, role, projectID, member, icon, banner, entries, links, tags)

	return err
}
