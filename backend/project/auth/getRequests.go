package auth

import (
	"context"
	"time"
)

//GetRequests retrieves the project join requests
func (auth *Auth) GetRequests(ctx context.Context, projectID, userID) error {
	//first check user permissions
	role, err := auth.FindUserRole(ctx, projectID, userID)
	if err != nil {
		return err
	}
	if !role.ManageMembers {
		return errInvalidPermission
	}


	//Add new member with given role
	err = auth.Strg.GetRequests(ctx, projectID)
	if err != nil {
		return err
	}

	return nil
}
