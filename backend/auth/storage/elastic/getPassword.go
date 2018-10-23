package elasticstore

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/olivere/elastic"
	"github.com/sea350/ustart_mono/backend/auth/authpb"
	"github.com/sea350/ustart_mono/backend/auth/storage"
)

//GetPassword retreivs a user's password
func (estor *ElasticStore) GetPassword(ctx context.Context, email string) (string, error) {
	//pull soted data attached to the email
	query := elastic.NewTermQuery("Email", strings.ToLower(email))
	res, err := estor.client.Search().
		Index(estor.eIndex).
		Query(query).
		Do(ctx)

	if err != nil {
		return "", err
	}
	//if there are no hits, then no one exists by that email
	if res.Hits.TotalHits < 1 {
		return "", storage.ErrUserDoesNotExist
	}

	//there should never be more than one result. If there is, there is an issue
	if res.Hits.TotalHits > 1 {
		return "", storage.ErrTooManyResults
	}
	var usr authpb.Stored

	for _, element := range res.Hits.Hits {
		err := json.Unmarshal(*element.Source, &usr)
		if err != nil {
			return "", err
		}
		break
	}

	return usr.Password, nil
}
