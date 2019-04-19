package profile

import (
	"context"

	"github.com/sea350/ustart_micro/backend/profile/profilepb"
)

// Lite retreives minimal profile data based off of a uuid
func (profile *Profile) Lite(ctx context.Context, req *profilepb.LiteRequest) (*profilepb.LiteResponse, error) {

	prof, err := profile.strg.Lookup(ctx, req.UUID)
	if err != nil {
		return nil, err
	}

	return &profilepb.LiteResponse{
		UUID:      prof.UUID,
		Username:  prof.Username,
		FirstName: prof.FirstName,
		LastName:  prof.LastName,
		Avatar:    prof.Avatar,
	}, nil

}
