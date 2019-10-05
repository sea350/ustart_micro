package auth

import (
	"context"
)

//RemoveRoleProfile checks permissions of the user and removes a role
func (auth *Auth) RemoveRoleProfile(ctx context.Context, projectID, userID, roleName string) error {
	//first check user permissions
	role, err := auth.FindUserRole(ctx, projectID, userID)
	if err != nil {
		return err
	}
	if !role.ManageMembers {
		return errInvalidPermission
	}

	//you cannot remove any default role
	if roleName == creator || roleName == moderator || roleName == member {
		return errDefaultProfile
	}

	//then check if role exists
	role, err = auth.Strg.GetRoleProfile(ctx, projectID, roleName)
	if err != nil {
		return err
	}

	//Check to make sure there is no member with that role
	mems, err := auth.Strg.GetMembers(ctx, projectID)
	if err != nil {
		return err
	}

	for _, m := range mems {
		if m.RoleName == roleName {
			return errInvalidInput
		}
	}

	//If you got here then all permission checkout, add the role
	err = auth.Strg.RemoveRole(ctx, projectID, roleName)

	return err
}
