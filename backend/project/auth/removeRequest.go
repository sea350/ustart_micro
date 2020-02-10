package auth

import (
	"context"
)

//RemoveRequest removes an applicant request, used to reject
func (auth *Auth) RemoveRequest(ctx context.Context, projectID, userID, removerID string) error {

	//retreive user role
	removerRole, err := auth.FindUserRole(ctx, projectID, userID)
	if err != nil {
		return err
	}

	//checks if user currently has permission to remove request, or if the user is removing their own request
	if removerID == userID || removerRole.ManageMembers {
		auth.Strg.RemoveRequest(ctx, projectID, removerID)
		if err != nil {
			return err
		}
		return nil
	}

	return errInvalidPermission

}
