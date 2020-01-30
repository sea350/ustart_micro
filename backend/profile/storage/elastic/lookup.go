package elasticstore

import (
	"context"
	"encoding/json"

	"github.com/olivere/elastic"
	"github.com/sea350/ustart_micro/backend/profile/profilepb"
)

// Lookup retreives a profile doc using a certain UUID
func (estor *ElasticStore) Lookup(ctx context.Context, uuid string) (profilepb.Profile, error) {
	var profile profilepb.Profile

	termQuery := elastic.NewTermQuery("UUID", uuid)
	res, err := estor.client.Search().
		Index(estor.eIndex).
		Query(termQuery).
		Do(ctx)

	if err != nil {
		return profile, err
	}

	// if there are no hits, then no one exists by that uuid
	if res.Hits.TotalHits < 1 {
		return profile, ErrUserDoesNotExist
	}

	// if theres more than a single result then a problem has occurred
	if res.Hits.TotalHits > 1 {
		return profile, ErrTooManyResults
	}

	for _, elem := range res.Hits.Hits {

		data, err := elem.Source.MarshalJSON()
		if err != nil {
			return profile, err
		}

		err = json.Unmarshal(data, &profile)
		return profile, err
	}

	return profile, nil
}
