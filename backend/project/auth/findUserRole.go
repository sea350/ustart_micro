package auth

import (
	"context"

	"github.com/sea350/ustart_micro/backend/project/projectpb"
)

//FindUserRole gets the role profile associated with a user
func (auth *Auth) FindUserRole(ctx context.Context, projectID, userID string) (projectpb.Role, error) {
	//find the role name of the user
	roleName, err := auth.Strg.GetMemberRole(ctx, userID, projectID)
	if err != nil {
		return projectpb.Role{}, err
	}

	//special case for "Creator" role
	if roleName == creator {
		return projectpb.Role{
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
	role, err := auth.Strg.GetRoleProfile(ctx, projectID, roleName)
	if err != nil {
		return projectpb.Role{}, err
	}

	return role, nil
}
