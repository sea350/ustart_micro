package auth

import (
	"context"
	"errors"
	"strings"

	"github.com/sea350/ustart_mono/backend/auth/authpb"

	elastic "github.com/olivere/elastic"
)

//Lookup checks elastic search if a document exists with the criteria of the req
func (eclient *ElasticAuth) Lookup(ctx context.Context, req *authpb.LookupRequest) (*authpb.LookupResponse, error) {
	//search for docs with the email specified in req
	termQuery := elastic.NewTermQuery("Email", strings.ToLower(req.Email))
	res, err := eclient.client.Search().
		Index(eIndex).
		Query(termQuery).
		Do(ctx)

	if err != nil {
		return &authpb.LookupResponse{}, err
	}

	//if there are no hits, then no one exists by that email
	if res.Hits.TotalHits < 1 {
		return &authpb.LookupResponse{}, nil
	}

	//if theres more than a single result then a problem has occurred
	if res.Hits.TotalHits > 1 {
		return &authpb.LookupResponse{}, errors.New("More than one result found")
	}
	return &authpb.LookupResponse{Exists: true}, nil
}
