package auth

import (
	"context"
	"time"
)

//AddMember checks permissions of the user and adds a new member
func (auth *Auth) AddMember(ctx context.Context, projectID, userID, newMemberID, roleName string) error {
	//first check user permissions
	role, err := auth.FindUserRole(ctx, projectID, userID)
	if err != nil {
		return err
	}
	if !role.ManageMembers {
		return errInvalidPermission
	}

	//then check if role exists
	_, err = auth.strg.GetRoleProfile(ctx, projectID, roleName)
	if err != nil {
		return err
	}

	//Add new member with given role
	err = auth.strg.AddMember(ctx, newMemberID, projectID, roleName, time.Now().Format(auth.timeFormat))
	if err != nil {
		return err
	}

	return nil
}
