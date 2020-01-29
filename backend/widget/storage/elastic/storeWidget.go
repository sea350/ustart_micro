package elasticstore

import (
	"context"

	"github.com/sea350/ustart_micro/backend/widget/widgetpb"
)

//StoreWidget creates a new ES document for storing a new widget
func (estor *Store) StoreWidget(ctx context.Context, index int, body string, time string, references []*widgetpb.Reference) error {

	_, err := estor.client.Index().
		Index(estor.eIndex).
		Type(estor.eType).
		BodyJson(widgetpb.Widget{
			Index:      int64(index),
			Body:       body,
			Time:       time,
			References: references,
		}).
		Do(ctx)
	return err

}
