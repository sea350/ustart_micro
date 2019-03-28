package auth

import "errors"

var (
	// ErrInvalidEmail Email is not in a valid format
	ErrInvalidEmail = errors.New("Email is not in a valid format")

	// ErrEmailInUse Email is already in use
	ErrEmailInUse = errors.New("Email is already in use")

	// ErrInvalidToken Token is invalid
	ErrInvalidToken = errors.New("Token is invalid")

	// ErrIncorrectPassword Password is invalid
	ErrIncorrectPassword = errors.New("Password is invalid")
)
