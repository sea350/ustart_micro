package elasticstore

import (
	"context"
	"strings"
	"github.com/olivere/elastic"
	"github.com/sea350/ustart_mono/backend/auth/storage"
)

// ChangePassword changes a user's password
func (estor *ElasticStore) ChangePassword(ctx context.Context, email string, newPassword string) error {
	//pull soted data attached to the email
	query := elastic.NewTermQuery("Email", strings.ToLower(email))
	res, err := estor.client.Search().
		Index(estor.eIndex).
		Query(query).
		Do(ctx)

	if err != nil {
		return err
	}
	//if there are no hits, then no one exists by that email
	if res.Hits.TotalHits < 1 {
		return storage.ErrUserDoesNotExist
	}

	//there should never be more than one result. If there is, there is an issue
	if res.Hits.TotalHits > 1 {
		return storage.ErrTooManyResults
	}

	var ID string

	for _, element := range res.Hits.Hits {
		ID = element.Id
		break
	}

	_, err = estor.client.Update().
		Index(estor.eIndex).
		Id(ID).
		Doc(map[string]interface{}{"Password": newPassword}).
		Do(ctx)

	return err
}
