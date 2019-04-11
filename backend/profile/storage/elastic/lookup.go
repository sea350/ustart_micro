package elasticstore

import (
	"context"
	"strings"

	"github.com/olivere/elastic"
)

// Lookup looks up if a document exists using a certain UUID
func (estor *ElasticStore) Lookup(ctx context.Context, uuid string) (string, error) {
	termQuery := elastic.NewTermQuery("UUID", strings.ToLower(uuid))
	res, err := estor.client.Search().
		Index(estor.eIndex).
		Query(termQuery).
		Do(ctx)

	if err != nil {
		return "", err
	}

	// if there are no hits, then no one exists by that uuid
	if res.Hits.TotalHits < 1 {
		return "", ErrUserDoesNotExist
	}

	// if theres more than a single result then a problem has occurred
	if res.Hits.TotalHits > 1 {
		return "", ErrTooManyResults
	}

	for _, elem := range res.Hits.Hits {
		return elem.Id, nil
	}
	return "", nil
}
