package profile

import "errors"

var (
	// ErrInvalidUsername Username is not in a valid format
	ErrInvalidUsername = errors.New("Username is not in a valid format")

	// ErrUsernameInUse Username is already in use
	ErrUsernameInUse = errors.New("Username is already in use")

	//ErrProfileExists This user already has a profile
	ErrProfileExists = errors.New("This user already has a profile")

	//ErrProblemLoadingProfile This user already has a profile
	ErrProblemLoadingProfile = errors.New("There was a problem loading one or more profiles in this list")

	//ErrNoProjectFound couldnt find project
	ErrNoProjectFound = errors.New("Could not find a project with that ID on display")
)
