package elasticstore

import "errors"

var (
	// ErrUserDoesNotExist user doesnt exist
	ErrUserDoesNotExist = errors.New("User does not exist")

	// ErrTooManyResults if more than one result per email shows
	ErrTooManyResults = errors.New("Too many results, a crititcal error has occurred")

	// ErrEmailInUse happens when a signup using an existing email occurs
	ErrEmailInUse = errors.New("Email in use")
)

//ErrUserDoesNotExist returns a standardized error
func (estor *ElasticStore) ErrUserDoesNotExist() error {
	return ErrUserDoesNotExist
}

//ErrTooManyResults returns a standardized error
func (estor *ElasticStore) ErrTooManyResults() error {
	return ErrTooManyResults
}

//ErrEmailInUse returns a standardized error
func (estor *ElasticStore) ErrEmailInUse() error {
	return ErrEmailInUse
}
