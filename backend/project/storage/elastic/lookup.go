package elasticstore

import (
	"context"
	"encoding/json"

	"github.com/olivere/elastic"
	"github.com/sea350/ustart_micro/backend/project/projectpb"
)

// Lookup retreives a project doc using a certain PID
func (estor *ElasticStore) Lookup(ctx context.Context, pid string) (projectpb.Project, error) {
	var project projectpb.Project

	termQuery := elastic.NewTermQuery("PID", pid)
	res, err := estor.client.Search().
		Index(estor.eIndex).
		Query(termQuery).
		Do(ctx)

	if err != nil {
		return project, err
	}

	// if there are no hits, then no one exists by that pid
	if res.Hits.TotalHits < 1 {
		return project, ErrProjectDoesNotExist
	}

	// if theres more than a single result then a problem has occurred
	if res.Hits.TotalHits > 1 {
		return project, ErrTooManyResults
	}

	for _, elem := range res.Hits.Hits {

		data, err := elem.Source.MarshalJSON()
		if err != nil {
			return project, err
		}

		err = json.Unmarshal(data, &project)
		return project, err
	}

	return project, nil
}
