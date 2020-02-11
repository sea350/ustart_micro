package elasticstore

import (
	"context"

	"github.com/sea350/ustart_micro/backend/widget/widgetpb"
)

//StoreWidget creates a new ES document for storing a new widget
func (estor *Store) StoreWidget(ctx context.Context, index int, body string, time string, references []*widgetpb.Reference) (string, error) {

	res, err := estor.client.Index().
		Index(estor.eIndex).
		BodyJson(widgetpb.Widget{
			Index:      int64(index),
			Body:       body,
			Time:       time,
			References: references,
		}).
		Do(ctx)

	if err != nil {
		return "", err
	}

	return res.Id, nil
}
