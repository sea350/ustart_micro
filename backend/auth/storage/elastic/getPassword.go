package elasticstore

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/olivere/elastic"
	"github.com/sea350/ustart_micro/backend/auth/authpb"
)

// GetPassword retreivs a user's password
func (estor *ElasticStore) GetPassword(ctx context.Context, email string) (string, error) {
	// pull sorted data attached to the email
	query := elastic.NewTermQuery("Email", strings.ToLower(email))
	res, err := estor.client.Search().
		Index(estor.eIndex).
		Query(query).
		Do(ctx)

	if err != nil {
		return "", err
	}

	// if there are no hits, then no one exists by that email
	if res.Hits.TotalHits < 1 {
		return "", ErrUserDoesNotExist
	}

	// there should never be more than one result. If there is, there is an issue
	if res.Hits.TotalHits > 1 {
		return "", ErrTooManyResults
	}

	var usr authpb.User
	err = json.Unmarshal(*res.Hits.Hits[0].Source, &usr)
	if err != nil {
		return "", err
	}

	return usr.Password, nil
}
