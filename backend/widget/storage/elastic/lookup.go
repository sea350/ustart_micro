package elasticstore

import (
	"context"
	"encoding/json"

	"github.com/sea350/ustart_micro/backend/widget/widgetpb"
)

//Lookup pulls a specific widget
func (estor *Store) Lookup(ctx context.Context, widgetID string) (widgetpb.Widget, error) {

	var wid widgetpb.Widget

	res, err := estor.client.Get().
		Index(estor.eIndex).
		Id(widgetID).
		Do(ctx)

	if err != nil {
		return wid, err
	}

	err = json.Unmarshal(*res.Source, &wid)

	return wid, err
}
