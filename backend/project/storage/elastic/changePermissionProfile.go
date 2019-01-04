package elasticstore

import (
	"context"

	"github.com/olivere/elastic"
)

//ChangePermissionProfile Changes a user's permission profile
func (estor *ElasticStore) ChangePermissionProfile(ctx context.Context, userID string, targerUserID string, projectID string) error {
	//Make sure user changing permission profile has correct permissions
	// query := elastic.NewTermQuery("UUID", userID)
	query := elastic.NewTermsQuery()
	res, err := estor.client.Search().
		Index(estor.eIndex).
		Query(query).
		Do(ctx)

	if err != nil {
		return err
	}

	//Pull target user from database
	//Update permission profile of targer user

}
