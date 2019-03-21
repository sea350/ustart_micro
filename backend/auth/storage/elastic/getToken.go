package elasticstore

import (
	"context"
	"errors"
	"time"
)

// GetToken looks up the token and expiration date associated with the given email
func (estor *ElasticStore) GetToken(ctx context.Context, email string) (string, time.Time, error) {

	return "", time.Time{}, errors.New("Feature has not been implemented yet")
}
