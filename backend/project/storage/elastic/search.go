package elasticstore

import (
	"context"
	"io"

	"github.com/olivere/elastic/v7"
)

// Search searches for... something... i think
func (estor *ElasticStore) Search(ctx context.Context, searchTerms []string, searchName, searchCustomURL, searchTags bool, mustMap map[string][]string, scrollID string) ([]string, error) {
	query := elastic.NewBoolQuery()
	// query = query.MustNot(elastic.NewTermQuery("Visible", false))
	// query = query.MustNot(elastic.NewTermQuery("Available", false))

	for _, term := range searchTerms {
		if searchName {
			query = query.Should(elastic.NewTermQuery("Name.keyword", term))

		}
		if searchCustomURL {
			query = query.Should(elastic.NewTermQuery("CustomURL.keyword", term))
		}
		if searchTags {
			query = query.Should(elastic.NewMatchQuery("Tags", term)) //.Fuzziness("AUTO"))
		}
	}

	arry, exists := mustMap["Category"]
	if exists {
		for _, term := range arry {
			query = query.Must(elastic.NewMatchQuery("Category", term))
		}
	}
	arry, exists = mustMap["Tags"]
	if exists {
		for _, term := range arry {
			query = query.Must(elastic.NewMatchQuery("Tags", term))
		}
	}

	results := []string{}

	scroll := estor.client.Scroll().
		Index(estor.eIndex).
		Query(query).
		Size(10).
		Sort("_score", false)

	if scrollID != "" {
		scroll = scroll.ScrollId(scrollID)
	}

	res, err := scroll.Do(ctx)
	if (err != nil && err != io.EOF) || res == nil {
		return results, err
	}

	for _, element := range res.Hits.Hits {
		results = append(results, element.Id)
	}
	return results, nil
}
