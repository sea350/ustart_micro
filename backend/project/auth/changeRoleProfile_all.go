package auth

import (
	"context"
	"errors"
)

//changeRoleProfile checks several conditions before modifying a role profile
func (auth *Auth) changeRoleProfile(ctx context.Context, projectID, userID, role, permission string, newValue bool) error {

	//Make sure the role exists
	changedRole, err := auth.Strg.GetRoleProfile(ctx, projectID, role)
	if err != nil {
		return err
	}
	//cannot modify creator profile
	if changedRole.Name == creator {
		return errInvalidPermission
	}

	//Now attempting to weed out potential permission conflicts
	changerRole, err := auth.FindUserRole(ctx, projectID, userID)
	if err != nil {
		return err
	}

	//changer needs to have permission
	if !changerRole.ManageMembers {
		return errInvalidPermission
	}

	//if you got here then the permissions check out
	//change the profile based on action
	switch permission {
	case "mem":
		err = auth.Strg.ModifyRoleManageMembers(ctx, projectID, role, newValue) //pass projectid, role name, new value | get error
	case "ico":
		err = auth.Strg.ModifyRoleChangeIcon(ctx, projectID, role, newValue)
	case "ban":
		err = auth.Strg.ModifyRoleChangeBanner(ctx, projectID, role, newValue)
	case "ent":
		err = auth.Strg.ModifyRoleManageEntries(ctx, projectID, role, newValue)
	case "lin":
		err = auth.Strg.ModifyRoleManageLinks(ctx, projectID, role, newValue)
	case "tag":
		err = auth.Strg.ModifyRoleManageTags(ctx, projectID, role, newValue)
	default:
		err = errors.New("Change role switch error: this shouldn't happen")
	}

	return err
}

//ChangeRoleManageMembers enters the proper permission for the classified action
//It's like this to avoid improper string input
func (auth *Auth) ChangeRoleManageMembers(ctx context.Context, projectID, userID, role string, newValue bool) error {
	return auth.changeRoleProfile(ctx, projectID, userID, role, "mem", newValue)
}

//ChangeRoleChangeIcon enters the proper permission for the classified action
//It's like this to avoid improper string input
func (auth *Auth) ChangeRoleChangeIcon(ctx context.Context, projectID, userID, role string, newValue bool) error {
	return auth.changeRoleProfile(ctx, projectID, userID, role, "ico", newValue)
}

//ChangeRoleChangeBanner enters the proper permission for the classified action
//It's like this to avoid improper string input
func (auth *Auth) ChangeRoleChangeBanner(ctx context.Context, projectID, userID, role string, newValue bool) error {
	return auth.changeRoleProfile(ctx, projectID, userID, role, "ban", newValue)
}

//ChangeRoleManageEntries enters the proper permission for the classified action
//It's like this to avoid improper string input
func (auth *Auth) ChangeRoleManageEntries(ctx context.Context, projectID, userID, role string, newValue bool) error {
	return auth.changeRoleProfile(ctx, projectID, userID, role, "man", newValue)
}

//ChangeRoleManageLinks enters the proper permission for the classified action
//It's like this to avoid improper string input
func (auth *Auth) ChangeRoleManageLinks(ctx context.Context, projectID, userID, role string, newValue bool) error {
	return auth.changeRoleProfile(ctx, projectID, userID, role, "lin", newValue)
}

//ChangeRoleManageTags enters the proper permission for the classified action
//It's like this to avoid improper string input
func (auth *Auth) ChangeRoleManageTags(ctx context.Context, projectID, userID, role string, newValue bool) error {
	return auth.changeRoleProfile(ctx, projectID, userID, role, "tag", newValue)
}
