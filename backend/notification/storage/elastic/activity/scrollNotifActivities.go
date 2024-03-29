package elasticstore

import (
	"context"
	"encoding/json"
	"io"
	"log"

	"github.com/olivere/elastic"
	notifpb "github.com/sea350/ustart_micro/backend/notification/notificationpb"
)

//ScrollNotifActivities scrolls through a set of activities, returns results and scroll id
func (estor *ElasticStore) ScrollNotifActivities(ctx context.Context, activityIDs []string, scrollID string) ([]notifpb.Activity, string, error) {

	activQuery := elastic.NewBoolQuery()

	for _, activID := range activityIDs {
		activQuery = activQuery.Should(elastic.NewTermQuery("_id", activID))
	}

	mapResults := []notifpb.Activity{}

	scroll := estor.client.Scroll().
		Index(estor.eIndex).
		Query(activQuery).
		Sort("Timestamp", false).
		Size(10)

	if len(scrollID) > 0 {
		scroll = scroll.ScrollId(scrollID)
	}

	res, scrollErr := scroll.Do(ctx)

	if scrollErr != nil && scrollErr != io.EOF {
		return mapResults, "", scrollErr
	}

	if res.TotalHits() == 0 {
		return mapResults, "", nil
	}

	var activ notifpb.Activity

	for _, hit := range res.Hits.Hits {
		if scrollErr == io.EOF {
			continue
		}
		err := json.Unmarshal(*hit.Source, &activ)
		if err != nil && err != io.EOF {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println(err)
			continue
		}

		activ.ActivityID = hit.Id

		mapResults = append(mapResults, activ)

	}

	return mapResults, res.ScrollId, nil
}
