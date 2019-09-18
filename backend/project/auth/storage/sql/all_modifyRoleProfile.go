package sqlstore

import (
	"context"
	"fmt"
)

//ModifyRoleProfile changes a single permission in a role
func (dbConn *SQLStore) ModifyRoleProfile(ctx context.Context, uuid, projectID, role string, newValue bool) error {
	queryString := fmt.Sprintf(
		`UPDATE %s
		SET %s = $3
		WHERE uuid = $1 AND project_id = $2;`,
		dbConn.roleTN, role)

	_, err := dbConn.db.ExecContext(ctx, queryString, uuid, projectID, newValue)

	return err
}

//ModifyRoleManageMembers changes the manage member permission in a role
func (dbConn *SQLStore) ModifyRoleManageMembers(ctx context.Context, uuid, projectID string, newValue bool) error {

	err := dbConn.ModifyRoleProfile(ctx, uuid, projectID, "manage_members", newValue)

	return err
}

//ModifyRoleChangeIcon changes the change icon permission in a role
func (dbConn *SQLStore) ModifyRoleChangeIcon(ctx context.Context, uuid, projectID string, newValue bool) error {

	err := dbConn.ModifyRoleProfile(ctx, uuid, projectID, "change_icon", newValue)

	return err
}

//ModifyRoleChangeBanner changes the change banner permission in a role
func (dbConn *SQLStore) ModifyRoleChangeBanner(ctx context.Context, uuid, projectID string, newValue bool) error {

	err := dbConn.ModifyRoleProfile(ctx, uuid, projectID, "change_banner", newValue)

	return err
}

//ModifyRoleManageEntries changes the manage entries permission in a role
func (dbConn *SQLStore) ModifyRoleManageEntries(ctx context.Context, uuid, projectID string, newValue bool) error {

	err := dbConn.ModifyRoleProfile(ctx, uuid, projectID, "manage_entries", newValue)

	return err
}

//ModifyRoleManageLinks changes the manage links permission in a role
func (dbConn *SQLStore) ModifyRoleManageLinks(ctx context.Context, uuid, projectID string, newValue bool) error {

	err := dbConn.ModifyRoleProfile(ctx, uuid, projectID, "manage_links", newValue)

	return err
}

//ModifyRoleManageTags changes the manage tags permission in a role
func (dbConn *SQLStore) ModifyRoleManageTags(ctx context.Context, uuid, projectID string, newValue bool) error {

	err := dbConn.ModifyRoleProfile(ctx, uuid, projectID, "manage_tags", newValue)

	return err
}
