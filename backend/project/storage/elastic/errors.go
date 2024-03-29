package elasticstore

import "errors"

var (
	errProjectDoesNotExist = errors.New("Project does not exist")

	//ErrCustomURLInUse Project Custom URL is in use
	ErrCustomURLInUse = errors.New("Project Custom URL is in use")

	// ErrTooManyResults if more than one result per custom url shows
	ErrTooManyResults = errors.New("Too many results, a critical error has occurred")

	//ErrCustomURLDoesNotExist Url not found
	ErrCustomURLDoesNotExist = errors.New("Url not found")
)

//ErrProjectDoesNotExist project doesnt exist
func (e *ElasticStore) ErrProjectDoesNotExist() error {
	return errProjectDoesNotExist
}
