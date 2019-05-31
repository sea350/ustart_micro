package elasticstore

import "errors"

var (
	// errUserDoesNotExist user doesnt exist
	errUserDoesNotExist = errors.New("User does not exist")

	// errTooManyResults if more than one result per email shows
	errTooManyResults = errors.New("Too many results, a crititcal error has occurred")

	// errEmailInUse happens when a signup using an existing email occurs
	errEmailInUse = errors.New("Email in use")
)

//ErrUserDoesNotExist returns a standardized error
func (estor *ElasticStore) ErrUserDoesNotExist() error {
	return errUserDoesNotExist
}

//ErrTooManyResults returns a standardized error
func (estor *ElasticStore) ErrTooManyResults() error {
	return errTooManyResults
}

//ErrEmailInUse returns a standardized error
func (estor *ElasticStore) ErrEmailInUse() error {
	return errEmailInUse
}
