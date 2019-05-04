package elasticstore

import (
	"context"

	"github.com/sea350/ustart_micro/backend/profile/profilepb"
)

//Register creates a new ES document for a new registering profile
func (estor *ElasticStore) Register(ctx context.Context, uuid string, username string, firstname string, lastname string, avatar string, banner string, dob string, visible bool, available bool) error {

	//Lock just to make sure no two people can sign up with the same username at the same time
	newUserLock.Lock()
	defer newUserLock.Unlock()

	_, err := estor.client.Index().
		Index(estor.eIndex).
		Type(estor.eType).
		BodyJson(profilepb.Profile{
			UUID:      uuid,
			Username:  username,
			FirstName: firstname,
			LastName:  lastname,
			Avatar:    avatar,
			Banner:    banner,
			DOB:       dob,
			Available: available,
			Visible:   visible,
		}).
		Id(uuid).
		Do(ctx)

	return err
}
