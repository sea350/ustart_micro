package elastic

import (
	"context"
	//"github.com/sea350/ustart_mono/backend/widget/wigetspb"
)

//Register creates a new ES document for storing a new widget
func (estor *ElasticStore) StoreWidget(ctx context.Context, widgetID string) error {

	//Lock just to make sure no two people can sign up with the same email at the same time
	newWidgetLock.Lock()
	defer newWidgetLock.Unlock()

	exists, err := estor.client.IndexExists(estor.eIndex).Do(ctx)
	if err != nil {
		panic(err)
	}

	if !exists {
		createIndex, err := estor.client.CreateIndex(estor.eIndex).BodyString("").Do(ctx) //DONT FORGET TO ADD MAPPTING LATER
		if err != nil {
			panic(err)
		}
		// TODO fix this.
		if !createIndex.Acknowledged {
			panic(err)
		}
	}

	/*
		_, err = estor.client.Index().
			Index(estor.eIndex).
			Type(estor.eType).
			BodyJson(widgetspb.User{WidgetID: widgetID}).
			Do(ctx)
	*/
	//Widgetspb doesn't work, and when I import it it seems to delete itself???

}
