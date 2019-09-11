package project

import "errors"

var (
	// ErrInvalidCustomURL CustomURL is not in a valid format
	ErrInvalidCustomURL = errors.New("CustomURL is not in a valid format")

	// ErrCustomURLInUse CustomURL is already in use
	ErrCustomURLInUse = errors.New("CustomURL is already in use")

	// //ErrProjectExists ...
	// ErrProjectExists = errors.New("This user already has a profile")

	//ErrProblemLoadingProfile This user already has a profile
	ErrProblemLoadingProject = errors.New("There was a problem loading one or more projects in this list")
)
