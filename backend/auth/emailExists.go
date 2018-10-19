package auth

import (
	"context"
	"strings"

	elastic "github.com/olivere/elastic"
)

//EmailExists checks elastic search if the email entered is in use
func (eclient *Eclient) EmailExists(ctx context.Context, email string) bool {
	termQuery := elastic.NewTermQuery("Email", strings.ToLower(email))
	res, err := eclient.client.Search().
		Index(eIndex).
		Query(termQuery).
		Do(ctx)

	if err != nil {
		return false
	}

	if res.Hits.TotalHits > 0 {
		return true
	}
	return false
}
