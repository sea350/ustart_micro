package auth

import (
	"context"

	"github.com/sea350/ustart_micro/backend/project/projectpb"
)

//GetRoles retrieves the project role profiles
func (auth *Auth) GetRoles(ctx context.Context, projectID, userID string) ([]*projectpb.Role, error) {
	//first check user permissions
	role, err := auth.FindUserRole(ctx, projectID, userID)
	if err != nil {
		return []*projectpb.Role{}, err
	}
	if !role.ManageMembers {
		return []*projectpb.Role{}, errInvalidPermission
	}

	//Add new member with given role
	reqs, err := auth.Strg.GetProjectRoles(ctx, projectID)
	if err != nil {
		return []*projectpb.Role{}, err
	}

	return reqs, nil
}
