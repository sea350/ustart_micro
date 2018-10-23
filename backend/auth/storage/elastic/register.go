package elasticstore

import (
	"context"

	"github.com/sea350/ustart_mono/backend/auth/authpb"
	"github.com/sea350/ustart_mono/backend/auth/storage"
)

//Register creates a new ES document for a new registering user
func (estor *ElasticStore) Register(ctx context.Context, email string, password, token string, accountType string) error {

	//Lock just to make sure no two people can sign up with the same email at the same time
	newUserLock.Lock()
	defer newUserLock.Unlock()

	// make sure email is not in use
	exists, err := estor.Lookup(ctx, email)
	if err != nil {
		return err
	}
	if exists {
		return storage.ErrEmailInUse
	}

	//before instering into database make sure the index exists
	exists, err = estor.client.IndexExists(estor.eIndex).Do(ctx)
	if err != nil {
		panic(err)
	}

	// If the index doesn't exist, create it. there shouldnt be any errors but if there are they are critical
	if !exists {
		createIndex, err := estor.client.CreateIndex(estor.eIndex).BodyString("").Do(ctx) //DONT FORGET TO ADD MAPPTING LATER
		if err != nil {
			panic(err)
		}
		// TODO fix this.
		if !createIndex.Acknowledged {
			panic(err)
		}
	}

	store := authpb.Stored{Email: email, Password: password}

	_, err = estor.client.Index().
		Index(estor.eIndex).
		Type(estor.eType).
		BodyJson(store).
		Do(ctx)

	return err
}
