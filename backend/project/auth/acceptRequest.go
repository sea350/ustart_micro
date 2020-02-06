package auth

import (
	"context"
)

//AcceptRequest checks whether the user is currently part of the project, before removing them from the request list and dding them to the project as a member
func (auth *Auth) AcceptRequest(ctx context.Context, projectID, userID, newMemberID, newRole string) error {

	//first checks if user currently is part of the project
	role, err := auth.FindUserRole(ctx, projectID, userID)
	if err != nil {
		return err
	}

	if role.ManageMembers {
		err = auth.RemoveRequest(ctx, projectID, userID, newMemberID)
		if err != nil {
			return err
		}
		err = auth.AddMember(ctx, projectID, userID, newMemberID, newRole)
		if err != nil {
			return err
		}
	}

	return nil
}
