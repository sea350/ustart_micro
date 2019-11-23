package elasticstore

import (
	"context"

	notifpb "github.com/sea350/ustart_micro/backend/notification/notificationpb"
)

//NewNotif stores an notification in ES doc form
func (estor *ElasticStore) NewNotif(ctx context.Context, uuid, activityID, timestamp string) error {

	_, err := estor.client.Index().
		Index(estor.eIndex).
		Type(estor.eType).
		BodyJson(notifpb.Notification{
			UUID:       uuid,
			ActivityID: activityID,
			Timestamp:  timestamp,
		}).
		Do(ctx)

	return err
}
