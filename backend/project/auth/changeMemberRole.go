package auth

import (
	"context"
)

//ChangeMemberRole checks several conditions before modifying a member's role
func (auth *Auth) ChangeMemberRole(ctx context.Context, projectID, changerID, changedID, newRole string) error {

	//Make sure the role exists
	_, err := auth.strg.GetRoleProfile(ctx, projectID, newRole)
	if err != nil {
		return err
	}

	//Now attempting to weed out potential permission conflicts
	changerRole, err := auth.FindUserRole(ctx, projectID, changerID)
	if err != nil {
		return err
	}

	//changer needs to have permission
	if !changerRole.ManageMembers {
		return errInvalidPermission
	}
	//only creators can modify creators
	changedRole, err := auth.FindUserRole(ctx, projectID, changedID)
	if err != nil {
		return err
	}
	if changedRole.Name == creator && changerRole.Name != creator {
		return errInvalidPermission
	}

	//if you got here then the permissions check out
	//change the member

	err = auth.strg.ModifyMemberRole(ctx, changedID, projectID, newRole)

	return err
}
