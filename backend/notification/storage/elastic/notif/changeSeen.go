package elasticstore

import (
	"context"
)

//ChangeSeen modifies the seen value of a notification
func (estor *ElasticStore) ChangeSeen(ctx context.Context, notifID string, newValue bool) error {

	_, err := estor.client.Update().
		Index(estor.eIndex).
		Type(estor.eType).
		Id(notifID).
		Doc(map[string]interface{}{"Seen": newValue}).
		Do(ctx)

	return err
}
