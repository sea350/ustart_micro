package profile

import (
	"context"

	"github.com/sea350/ustart_micro/backend/profile/profilepb"
)

//ToggleAvailable turns a profile's availability on/off
func (profile *Profile) ToggleAvailable(ctx context.Context, req *profilepb.ToggleAvailableRequest) (*profilepb.ToggleAvailableResponse, error) {
	prof, err := profile.strg.Lookup(ctx, req.UUID)
	if err != nil {
		return nil, err
	}

	if prof.Available {
		err = profile.strg.UpdateAvailable(ctx, req.UUID, false)
	} else {
		err = profile.strg.UpdateAvailable(ctx, req.UUID, true)
	}
	if err != nil {
		return nil, err
	}

	return &profilepb.ToggleAvailableResponse{NewValue: !prof.Available}, nil

}
