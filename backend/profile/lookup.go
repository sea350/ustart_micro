package profile

import (
	"context"

	"github.com/sea350/ustart_micro/backend/profile/profilepb"
)

// Lookup checks to see if a username is in use
func (profile *Profile) Lookup(ctx context.Context, req *profilepb.LookupRequest) (*profilepb.LookupResponse, error) {

	_, err := profile.strg.LookupUsername(ctx, req.Username)
	if err == profile.strg.ErrUserDoesNotExist() {
		return &profilepb.LookupResponse{Exists: false}, nil
	}
	if err != nil {
		return nil, err
	}

	return &profilepb.LookupResponse{Exists: true}, nil

}
