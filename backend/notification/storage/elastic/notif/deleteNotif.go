package elasticstore

import (
	"context"
)

//DeleteNotif removes a notification
func (estor *ElasticStore) DeleteNotif(ctx context.Context, notifID string) error {

	_, err := estor.client.Delete().
		Index(estor.eIndex).
		Type(estor.eType).
		Id(notifID).
		Do(ctx)

	return err
}
