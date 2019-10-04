package auth

import (
	"context"
	"strings"
)

//AddRoleProfile checks permissions of the user and adds a new role
func (auth *Auth) AddRoleProfile(ctx context.Context, projectID, userID, roleName string, member, icon, banner, entries, links, tags bool) error {
	//first check user permissions
	role, err := auth.FindUserRole(ctx, projectID, userID)
	if err != nil {
		return err
	}
	if !role.ManageMembers {
		return errInvalidPermission
	}

	//then check if role exists
	roles, err := auth.strg.GetProjectRoles(ctx, projectID)
	if err != nil {
		return err
	}

	for _, role := range roles {
		//Make the casing equivalent to try to avoid shady new roles
		//This will probably need to be made more sophisticated later
		if strings.ToLower(role.Name) == strings.ToLower(roleName) {
			return errInvalidInput
		}
	}

	//If you got here then all permission checkout, add the role
	err = auth.strg.AddRole(ctx, roleName, projectID, member, icon, banner, entries, links, tags)

	return err
}
