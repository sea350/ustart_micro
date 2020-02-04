package elasticstore

import (
	"context"
)

//update for local use only
func (estor *ElasticStore) update(ctx context.Context, uuid, field string, newValue interface{}) error {

	_, err := estor.client.Update().
		Index(estor.eIndex).
		Type(estor.eType).
		Id(uuid).
		Doc(map[string]interface{}{field: newValue}).
		Do(ctx)

	return err
}

//UpdateBasicInfo updates listed fields
func (estor *ElasticStore) UpdateBasicInfo(ctx context.Context, uuid, firstName, lastname, desc, org, dob string) error {
	var err error
	err = estor.update(ctx, uuid, "FirstName", firstName)
	err = estor.update(ctx, uuid, "LastName", lastname)
	err = estor.update(ctx, uuid, "Description", desc)
	err = estor.update(ctx, uuid, "Organization", org)
	err = estor.update(ctx, uuid, "DOB", dob)
	return err
}

//UpdateAvatar updates that field
func (estor *ElasticStore) UpdateAvatar(ctx context.Context, uuid, newImgLink string) error {
	err := estor.update(ctx, uuid, "Avatar", newImgLink)
	return err
}

//UpdateBanner updates that field
func (estor *ElasticStore) UpdateBanner(ctx context.Context, uuid, newImgLink string) error {
	err := estor.update(ctx, uuid, "Banner", newImgLink)
	return err
}

//UpdateAvailable updates that field
func (estor *ElasticStore) UpdateAvailable(ctx context.Context, uuid string, available bool) error {
	err := estor.update(ctx, uuid, "Available", available)
	return err
}
