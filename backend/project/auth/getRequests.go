package auth

import (
	"context"
)

//GetRequests retrieves the project join requests
func (auth *Auth) GetRequests(ctx context.Context, projectID, userID string) ([]string, error) {
	//first check user permissions
	role, err := auth.FindUserRole(ctx, projectID, userID)
	if err != nil {
		return []string{}, err
	}
	if !role.ManageMembers {
		return []string{}, errInvalidPermission
	}

	//Add new member with given role
	reqs, err := auth.Strg.GetRequests(ctx, projectID)
	if err != nil {
		return []string{}, err
	}

	return reqs, nil
}
