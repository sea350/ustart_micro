package elasticstore

import "errors"

var (
	errNoNotifsFound = errors.New("No notifications found")

	errNotifNotFound = errors.New("Notification not found")
)

//ErrNoNotifsFound returns a standardized error
func (estor *ElasticStore) ErrNoNotifsFound() error {
	return errNoNotifsFound
}

//ErrNotifNotFound returns a standardized error
func (estor *ElasticStore) ErrNotifNotFound() error {
	return errNotifNotFound
}
