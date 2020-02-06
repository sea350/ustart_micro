package auth

import (
	"context"
)

//RemoveRequest checks whether the user is currently part of the project and sends the request to the queue if not
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
