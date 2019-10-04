package auth

import (
	"context"

	"github.com/sea350/ustart_micro/backend/project/auth/types"
)

//FindUserRole gets the role profile associated with a user
func (auth *Auth) FindUserRole(ctx context.Context, projectID, userID string) (types.Role, error) {
	//find the role name of the user
	roleName, err := auth.strg.GetMemberRole(ctx, userID, projectID)
	if err != nil {
		return types.Role{}, err
	}

	//special case for "Creator" role
	if roleName == creator {
		return types.Role{
			Name:          creator,
			ManageMembers: true,
			ChangeIcon:    true,
			ChangeBanner:  true,
			ManageEntries: true,
			ManageLinks:   true,
			ManageTags:    true,
		}, err
	}

	//pull the role profile associated with the name
	role, err := auth.strg.GetRoleProfile(ctx, projectID, roleName)
	if err != nil {
		return types.Role{}, err
	}

	return role, nil
}
