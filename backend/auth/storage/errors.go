package storage

import "errors"

var (
	// ErrUserDoesNotExist user doesnt exist
	ErrUserDoesNotExist = errors.New("User does not exist")

	// ErrTooManyResults if more than one result per email shows
	ErrTooManyResults = errors.New("Too many results, a crititcal error has occurred")

	// ErrEmailInUse happens when a signup using an existing email occurs
	ErrEmailInUse = errors.New("Email in use")
)
