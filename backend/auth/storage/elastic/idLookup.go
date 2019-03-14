package elasticstore

import (
	"context"
	"errors"
)

// IDLookup looks up if a document exists using a certain uuid
func (estor *ElasticStore) IDLookup(ctx context.Context, uuid string) (bool, error) {

	return false, errors.New("Feature has not been implemented yet")
}
