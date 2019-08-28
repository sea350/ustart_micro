package sqlstore

import "errors"

var (
	// errSessionDoesNotExist session not found
	errSessionDoesNotExist = errors.New("Session does not exist")
)

//ErrSessionDoesNotExist returns a standardized error
func (dbConn *SQLStore) ErrSessionDoesNotExist() error {
	return errSessionDoesNotExist
}
