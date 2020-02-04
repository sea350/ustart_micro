package elasticstore

import (
	"context"
)

//update for local use only
func (estor *ElasticStore) update(ctx context.Context, uuid, field, newValue string) error {

	_, err := estor.client.Update().
		Index(estor.eIndex).
		Type(estor.eType).
		Id(uuid).
		Doc(map[string]interface{}{field: newValue}).
		Do(ctx)

	return err
}

//UpdateBasicInfo updates that field
func (estor *ElasticStore) UpdateBasicInfo(ctx context.Context, uuid, firstName, lastname, desc, org, dob string) error {
	var err error
	err = estor.update(ctx, uuid, "FirstName", firstName)
	err = estor.update(ctx, uuid, "LastName", lastname)
	err = estor.update(ctx, uuid, "Description", desc)
	err = estor.update(ctx, uuid, "Organization", org)
	err = estor.update(ctx, uuid, "DOB", dob)
	return err
}
