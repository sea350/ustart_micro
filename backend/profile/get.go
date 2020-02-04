package profile

import (
	"context"

	"github.com/sea350/ustart_micro/backend/profile/profilepb"
)

// Get retreives all profile data based off of a uuid
func (profile *Profile) Get(ctx context.Context, req *profilepb.GetRequest) (*profilepb.GetResponse, error) {

	prof, err := profile.strg.Lookup(ctx, req.UUID)
	if err != nil {
		return nil, err
	}

	return &profilepb.GetResponse{UserProfile: &prof}, nil

}
