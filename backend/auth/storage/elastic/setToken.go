package elasticstore

import (
	"context"
	"errors"
	"time"
)

// SetToken sets the token and expiration for the acc of the given email
func (estor *ElasticStore) SetToken(ctx context.Context, email string, token string, expiration time.Time) error {

	return errors.New("Feature has not been implemented yet")
}
