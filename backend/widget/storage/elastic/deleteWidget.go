package elasticstore

import (
	"context"
	//"github.com/sea350/ustart_micro/backend/widget/wigetspb" --> Doesn't want to implement???
)

//DeleteWidget ...
func (estor *Store) DeleteWidget(ctx context.Context, widgetID string) error {
	_, err := estor.client.Delete().
		Index(estor.eIndex).
		Type(estor.eType).
		Id(widgetID).
		Do(ctx)

	return err
}
