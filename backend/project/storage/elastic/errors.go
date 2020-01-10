package elasticstore

import "errors"

var (
	// ErrProjectDoesNotExist project doesnt exist
	ErrProjectDoesNotExist = errors.New("Project does not exist")

	// ErrProjectDoesNotExist project doesnt exist
	ErrCustomURLInUse = errors.New("Project Custom URL is in use")

	// ErrTooManyResults if more than one result per custom url shows
	ErrTooManyResults = errors.New("Too many results, a critical error has occurred")

//
)
