package auth

import (
	"context"
	"strings"

	"github.com/sea350/ustart_mono/backend/auth/authpb"

	elastic "github.com/olivere/elastic"
)

//Lookup checks elastic search if a document exists with the criteria of the req
func (eclient *Eclient) Lookup(ctx context.Context, req *authpb.LookupRequest) (*authpb.LookupResponse, error) {
	termQuery := elastic.NewTermQuery("Email", strings.ToLower(req.Email))
	res, err := eclient.client.Search().
		Index(eIndex).
		Query(termQuery).
		Do(ctx)

	if err != nil {
		return &authpb.LookupResponse{}, err
	}

	if res.Hits.TotalHits > 0 {
		return &authpb.LookupResponse{}, nil
	}
	return &authpb.LookupResponse{Exists: true}, nil
}
