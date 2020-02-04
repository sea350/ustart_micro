package elasticstore

import (
	"context"
	"encoding/json"
	"io"
	"log"

	"github.com/sea350/ustart_micro/backend/widget/widgetpb"

	"github.com/olivere/elastic"
)

//GetAfter pulls all of one owner's showcase widgets after a certain index
func (estor *Store) GetAfter(ctx context.Context, owner string, index int) (int, []widgetpb.Widget, error) {
	query := elastic.NewBoolQuery()
	query = query.Must(elastic.NewMatchQuery("References.ReferenceID", owner))
	query = query.Must(elastic.NewMatchQuery("References.Classification", "owner"))
	query = query.Must(elastic.NewRangeQuery("Index").Gt(index))
	results := []widgetpb.Widget{}

	search := estor.client.Search().
		Index(estor.eIndex).
		Query(query).
		Sort("Index", true)

	res, err := search.Do(ctx)
	if (err != nil && err != io.EOF) || res == nil {
		return 0, results, err
	}

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	for _, element := range res.Hits.Hits {
		data := widgetpb.Widget{}
		err = json.Unmarshal(*element.Source, &data)
		if err != nil {
			log.Println("Owner ID: ", owner, "| err: ", err)
			continue
		}
		data.WidgetID = element.Id
		results = append(results, data)
	}

	return int(res.Hits.TotalHits), results, nil
}
