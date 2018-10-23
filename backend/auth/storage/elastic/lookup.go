package elasticstore

import (
	"context"
	"strings"

	"github.com/olivere/elastic"
	"github.com/sea350/ustart_mono/backend/auth/storage"
)

// Lookup looks up if a document exists using a certain email
func (estor *ElasticStore) Lookup(ctx context.Context, email string) (bool, error) {

	termQuery := elastic.NewTermQuery("Email", strings.ToLower(email))
	res, err := estor.client.Search().
		Index(estor.eIndex).
		Query(termQuery).
		Do(ctx)

	if err != nil {
		return false, err
	}

	//if there are no hits, then no one exists by that email
	if res.Hits.TotalHits < 1 {
		return false, nil
	}

	//if theres more than a single result then a problem has occurred
	if res.Hits.TotalHits > 1 {
		return false, storage.ErrTooManyResults
	}
	return true, nil

}
