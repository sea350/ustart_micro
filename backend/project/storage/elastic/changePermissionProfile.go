package elasticstore

import (
	"context"

	"github.com/olivere/elastic"
)

//------------------------------------------WIP---------------------------------------------

//ChangePermissionProfile Changes a user's permission profile
func (estor *ElasticStore) ChangePermissionProfile(ctx context.Context, userID string, targetUserID string, projectID string) error {
	//Make sure user changing permission profile has correct permissions
	// query := elastic.NewTermQuery("UUID", userID)
	query := elastic.NewBoolQuery()
	query = query.Must(elastic.NewTermQuery("UUID", userID))
	query = query.Must(elastic.NewTermQuery("PID", projectID))
	res, err := estor.client.Search().
		Index(estor.eIndex).
		Query(query).
		Do(ctx)

	if err != nil {
		return err
	}

	//If no results, current user does not have permission to change permission or user not found
	if res.Hits.TotalHits < 1 {
		return ErrUserDoesNotExist
	}

	//If too many results, critical error. Should not be more than 1 result
	if res.Hits.TotalHits > 1 {
		return ErrTooManyResults
	}

	//Pull target user from database
	userQuery := elastic.NewBoolQuery()
	userQuery = userQuery.Must(elastic.NewTermQuery("UUID", targetUserID))

	userRes, err := estor.client.Search().
		Index(estor.eIndex).
		Query(query).
		Do(ctx)

	if err != nil {
		return err
	}

	//If no results, target user could not be found
	if userRes.Hits.TotalHits < 1 {
		return ErrUserDoesNotExist
	}

	//If too many results, critical error. Should not be more than 1 result
	if userRes.Hits.TotalHits > 1 {
		return ErrTooManyResults
	}

	//Update permission profile of targer user

}
