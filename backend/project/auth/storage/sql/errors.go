package sqlstore

import "errors"

var (
	// errNoResultsFound user doesnt exist
	errNoResultsFound = errors.New("No results found")

	// errTooManyResults if more than one result per email shows
	errTooManyResults = errors.New("Too many results, a crititcal error has occurred")
)

//ErrNoResultsFound returns a standardized error
func (dbConn *SQLStore) ErrNoResultsFound() error {
	return errNoResultsFound
}

//ErrTooManyResults returns a standardized error
func (dbConn *SQLStore) ErrTooManyResults() error {
	return errTooManyResults
}
