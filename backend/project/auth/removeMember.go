package auth

import (
	"context"
)

//RemoveMember checks several conditions before removing a member
func (auth *Auth) RemoveMember(ctx context.Context, projectID, removerID, removedID string) error {

	removerRole, err := auth.FindUserRole(ctx, projectID, removerID)
	if err != nil {
		return err
	}

	//Now attempting to weed out potential permission conflicts

	//first check user is removing self
	if removerID == removedID {
		//if user is creator, make sure there is more than one creator
		if removerRole.Name == creator {
			mems, err := auth.Strg.GetMembers(ctx, projectID)
			if err != nil {
				return err
			}
			var creatorCounter int
			for _, member := range mems {
				if member.RoleName == creator {
					creatorCounter++
				}
			}
			if creatorCounter <= 1 {
				return errInvalidPermission
			}
		}
	} else {
		//otherwise remover needs to have permission
		if !removerRole.ManageMembers {
			return errInvalidPermission
		}
		//only creators can remove creators
		removedRole, err := auth.FindUserRole(ctx, projectID, removedID)
		if err != nil {
			return err
		}
		if removedRole.Name == creator && removerRole.Name != creator {
			return errInvalidPermission
		}
	}

	//if you got here then the permissions check out
	//remove the member

	err = auth.Strg.RemoveMember(ctx, projectID, removedID)

	return err
}
