package elasticstore

import "errors"

var (
	errActivityNotFound = errors.New("There was a problem retreiving the activity")
)

//ErrActivityNotFound returns a standardized error
func (estor *ElasticStore) ErrActivityNotFound() error {
	return errActivityNotFound
}
