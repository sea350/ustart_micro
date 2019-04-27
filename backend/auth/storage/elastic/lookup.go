package elasticstore

import (
	"context"
	"strings"

	"github.com/olivere/elastic"
)

// Lookup looks up if a document exists using a certain email
func (estor *ElasticStore) Lookup(ctx context.Context, email string) (string, error) {
	termQuery := elastic.NewTermQuery("Email", strings.ToLower(email))
	res, err := estor.client.Search().
		Index(estor.eIndex).
		Query(termQuery).
		Do(ctx)

	if err != nil {
		return "", err
	}

	// if there are no hits, then no one exists by that email
	if res.Hits.TotalHits.Value < 1 {
		return "", nil
	}

	// if theres more than a single result then a problem has occurred
	if res.Hits.TotalHits.Value > 1 {
		return "", ErrTooManyResults
	}

	for _, elem := range res.Hits.Hits {
		return elem.Id, nil
	}
	return "", nil
}
