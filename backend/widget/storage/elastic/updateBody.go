package elasticstore

import (
	"context"
)

//UpdateBody overwrites a particular widget's content
func (estor *Store) UpdateBody(ctx context.Context, widgetID string, newBody string) error {

	_, err := estor.client.Update().
		Index(estor.eIndex).
		Id(widgetID).
		Doc(map[string]interface{}{"Body": newBody}).
		Do(ctx)

	return err
}
