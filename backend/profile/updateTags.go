package profile

import (
	"context"

	"github.com/sea350/ustart_micro/backend/profile/profilepb"
)

//UpdateTags updates a profile's banner
func (profile *Profile) UpdateTags(ctx context.Context, req *profilepb.UpdateTagsRequest) (*profilepb.UpdateTagsResponse, error) {

	err := profile.strg.UpdateTags(ctx, req.UUID, req.Tags)
	if err != nil {
		return nil, err
	}
	return &profilepb.UpdateTagsResponse{}, nil

}
