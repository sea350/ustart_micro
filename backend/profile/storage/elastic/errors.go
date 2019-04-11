package elasticstore

import "errors"

var (
	// ErrUserDoesNotExist user doesnt exist
	ErrUserDoesNotExist = errors.New("User does not exist")

	// ErrTooManyResults if more than one result per email shows
	ErrTooManyResults = errors.New("Too many results, a crititcal error has occurred")
)

//ErrUserDoesNotExist returns a standardized error
func (estor *ElasticStore) ErrUserDoesNotExist() error {
	return ErrUserDoesNotExist
}

//ErrTooManyResults returns a standardized error
func (estor *ElasticStore) ErrTooManyResults() error {
	return ErrTooManyResults
}
