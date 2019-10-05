package auth

import "errors"

var (
	errInvalidPermission = errors.New("This user does not have sufficient permission")

	errInvalidInput = errors.New("The submission you entered is not allowed, check the spelling and try again")

	errPermissionProfileInUse = errors.New("The permission profile you are trying to remove is in use, please make sure there are no members with this role")

	errDefaultProfile = errors.New("This permission profile cannot be modified in this way")
)
