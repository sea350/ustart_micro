package auth

import (
	"context"
	"strings"

	"github.com/sea350/ustart_micro/backend/project/auth/types"
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
	roles, err := auth.Strg.GetProjectRoles(ctx, projectID)
	if err != nil {
		return err
	}

	//Creator role doesnt have a formal role profile so it must be added to the check
	//actual permissions arent relevant
	roles = append(roles, types.Role{Name: creator})
	for _, role := range roles {
		//Make the casing equivalent to try to avoid shady new roles
		//This will probably need to be made more sophisticated later
		if strings.ToLower(role.Name) == strings.ToLower(roleName) {
			return errInvalidInput
		}
	}

	//If you got here then all permission checkout, add the role
	err = auth.Strg.AddRole(ctx, roleName, projectID, member, icon, banner, entries, links, tags)

	return err
}
