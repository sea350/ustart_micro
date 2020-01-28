package elasticstore

import (
	"context"

	"github.com/sea350/ustart_micro/backend/widget/widgetspb"
)

//StoreWidget creates a new ES document for storing a new widget
func (estor *Store) StoreWidget(ctx context.Context, index int, body string, time string) error {

	_, err := estor.client.Index().
		Index(estor.eIndex).
		Type(estor.eType).
		BodyJson(widgetspb.Widget{
			Index: int64(index),
			Body:  body,
			Time:  time,
		}).
		Do(ctx)
	return err

}
