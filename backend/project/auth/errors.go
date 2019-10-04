package auth

import "errors"

var (
	errInvalidPermission = errors.New("This user does not have sufficient permission")

	errInvalidInput = errors.New("The submission you entered is not allowed, check the spelling and try again")
)
