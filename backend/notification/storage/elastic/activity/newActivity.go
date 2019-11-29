package elasticstore

import (
	"context"

	notifpb "github.com/sea350/ustart_micro/backend/notification/notificationpb"
)

//NewActivity performs an activity in ES doc form
func (estor *ElasticStore) NewActivity(ctx context.Context, uuid, objectID, action, timestamp string) (string, error) {

	ret, err := estor.client.Index().
		Index(estor.eIndex).
		Type(estor.eType).
		BodyJson(notifpb.Activity{
			ActorID:   uuid,
			ObjectID:  objectID,
			Action:    action,
			Timestamp: timestamp,
		}).
		Do(ctx)

	return ret.Id, err
}
