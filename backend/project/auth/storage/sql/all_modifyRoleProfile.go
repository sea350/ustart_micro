package sqlstore

import (
	"context"
	"fmt"
)

//modifyRoleProfile changes a single permission in a role, FOR INTERNAL USE ONLY
func (dbConn *SQLStore) modifyRoleProfile(ctx context.Context, projectID, roleName, permission string, newValue bool) error {
	queryString := fmt.Sprintf(
		`UPDATE %s
		SET %s = $3
		WHERE project_id = $1 AND role_name = $2;`,
		dbConn.roleTN, permission)

	_, err := dbConn.db.ExecContext(ctx, queryString, projectID, roleName, newValue)

	return err
}

//ModifyRoleManageMembers changes the manage member permission in a role
func (dbConn *SQLStore) ModifyRoleManageMembers(ctx context.Context, projectID, roleName string, newValue bool) error {

	err := dbConn.modifyRoleProfile(ctx, projectID, roleName, "manage_members", newValue)

	return err
}

//ModifyRoleChangeIcon changes the change icon permission in a role
func (dbConn *SQLStore) ModifyRoleChangeIcon(ctx context.Context, projectID, roleName string, newValue bool) error {

	err := dbConn.modifyRoleProfile(ctx, projectID, roleName, "change_icon", newValue)

	return err
}

//ModifyRoleChangeBanner changes the change banner permission in a role
func (dbConn *SQLStore) ModifyRoleChangeBanner(ctx context.Context, projectID, roleName string, newValue bool) error {

	err := dbConn.modifyRoleProfile(ctx, projectID, roleName, "change_banner", newValue)

	return err
}

//ModifyRoleManageEntries changes the manage entries permission in a role
func (dbConn *SQLStore) ModifyRoleManageEntries(ctx context.Context, projectID, roleName string, newValue bool) error {

	err := dbConn.modifyRoleProfile(ctx, projectID, roleName, "manage_entries", newValue)

	return err
}

//ModifyRoleManageLinks changes the manage links permission in a role
func (dbConn *SQLStore) ModifyRoleManageLinks(ctx context.Context, projectID, roleName string, newValue bool) error {

	err := dbConn.modifyRoleProfile(ctx, projectID, roleName, "manage_links", newValue)

	return err
}

//ModifyRoleManageTags changes the manage tags permission in a role
func (dbConn *SQLStore) ModifyRoleManageTags(ctx context.Context, projectID, roleName string, newValue bool) error {

	err := dbConn.modifyRoleProfile(ctx, projectID, roleName, "manage_tags", newValue)

	return err
}
