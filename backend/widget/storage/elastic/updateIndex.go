package elasticstore

import (
	"context"
)

//UpdateIndex reassignes a particular widgets's index
func (estor *Store) UpdateIndex(ctx context.Context, widgetID string, newIndex int) error {

	_, err := estor.client.Update().
		Index(estor.eIndex).
		Id(widgetID).
		Doc(map[string]interface{}{"Index": newIndex}).
		Do(ctx)

	return err
}
