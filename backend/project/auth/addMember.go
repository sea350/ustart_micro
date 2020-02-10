package auth

import (
	"context"
	"time"
)

//AddMember checks permissions of the user and adds a new member
func (auth *Auth) AddMember(ctx context.Context, projectID, userID, newMemberID, roleName string) error {
	memberMod.Lock()
	defer memberMod.Unlock()
	//check if theyre already a member
	_, err := auth.FindUserRole(ctx, projectID, newMemberID)
	if err != auth.Strg.ErrNoResultsFound() {
		if err != nil {
			return err
		}
		return errAlreadyExists
	}

	//first check user permissions
	role, err := auth.FindUserRole(ctx, projectID, userID)
	if err != nil {
		return err
	}
	if !role.ManageMembers {
		return errInvalidPermission
	}

	//then check if role exists, default to member if anything goes wrong
	_, err = auth.Strg.GetRoleProfile(ctx, projectID, roleName)
	if err != nil {
		roleName = "Member"
	}

	//Add new member with given role
	err = auth.Strg.AddMember(ctx, newMemberID, projectID, roleName, time.Now().Format(auth.timeFormat))
	if err != nil {
		return err
	}

	return nil
}
