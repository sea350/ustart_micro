package store

import (
	"context"

	"github.com/olivere/elastic"
)

//ChangePermissionProfileLocation Permission profile - Allows user to change project location
func (estor *ElasticStore) ChangePermissionProfileLocation(ctx context.Context, projectID string, profileName string, canChangeLocation bool) error {
	//Pull target project/permission profile from ES
	query := elastic.NewBoolQuery()
	query = query.Must(elastic.NewTermQuery("UUID", profileName))
	query = query.Must(elastic.NewTermQuery("PID", projectID))

	results, err := estor.client.Search().
		Index(estor.eIndex).
		Query(query).
		Do(ctx)

	if err != nil {
		return err
	}

	//If no results, project/profile could not be found
	if results.Hits.TotalHits < 1 {
		return ErrNoResults
	}

	//If too many results, critical error. Should not be more than 1 result
	if results.Hits.TotalHits > 1 {
		return ErrTooManyResults
	}

	//Update permission
	_, err = estor.client.Update().
		Index(estor.eIndex).
		Id(results.Hits.Hits[0].Id).
		Doc(map[string]interface{}{"CanChangeLocation": canChangeLocation}).
		Do(ctx)

	return err

}
