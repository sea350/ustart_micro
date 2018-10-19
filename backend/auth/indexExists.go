package auth

import (
	"context"
)

//IndexExists checks elastic search that the current index exists
func (eclient *Eclient) IndexExists(ctx context.Context) {
	exists, err := eclient.client.IndexExists(eIndex).Do(ctx)
	if err != nil {
		panic(err)
	}

	// If the index doesn't exist, create it and return error.
	if !exists {
		createIndex, err := eclient.client.CreateIndex(eIndex).BodyString("").Do(ctx)
		if err != nil {
			panic(err)
		}
		// TODO fix this.
		if !createIndex.Acknowledged {
			panic(err)
		}
	}
}
