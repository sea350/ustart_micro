package profile

import (
	"context"

	"github.com/sea350/ustart_micro/backend/profile/profilepb"
)

// Pull retreives all profile data based off of a username
func (profile *Profile) Pull(ctx context.Context, req *profilepb.PullRequest) (*profilepb.PullResponse, error) {

	prof, err := profile.strg.LookupUsername(ctx, req.Username)
	if err != nil {
		return nil, err
	}

	return &profilepb.PullResponse{UserProfile: &prof}, nil

}
