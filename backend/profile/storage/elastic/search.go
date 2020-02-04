package elasticstore

import (
	"context"
	"io"

	"github.com/olivere/elastic/v7"
)

// Search searches for... something... i think
func (estor *ElasticStore) Search(ctx context.Context, searchTerms []string, searchName, searchUsername, searchTags bool, mustMap map[string][]string, scrollID string) ([]string, error) {
	query := elastic.NewBoolQuery()
	query = query.MustNot(elastic.NewTermQuery("Visible", false))
	query = query.MustNot(elastic.NewTermQuery("Available", false))

	for _, term := range searchTerms {
		if searchName {
			query = query.Should(elastic.NewFuzzyQuery("FirstName", term).Fuzziness("AUTO"))
			query = query.Should(elastic.NewFuzzyQuery("LastName", term).Fuzziness("AUTO"))
		}
		if searchUsername {
			query = query.Should(elastic.NewFuzzyQuery("Username", term).Fuzziness("AUTO"))
		}
		if searchTags {
			query = query.Should(elastic.NewFuzzyQuery("Tags", term).Fuzziness("AUTO"))
		}
	}

	arry, exists := mustMap["Major"]
	if exists {
		for _, term := range arry {
			query = query.Must(elastic.NewMatchQuery("Majors", term))
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
