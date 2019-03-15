package auth

import "errors"

var (
	// ErrInvalidEmail Email is not in a valid format
	ErrInvalidEmail = errors.New("Email is not in a valid format")

	// ErrEmailInUse Email is already in use
	ErrEmailInUse = errors.New("Email is already in use")
)
