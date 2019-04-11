package profile

import (
	"context"

	"github.com/sea350/ustart_micro/backend/profile/profilepb"
)

// Register is a generic register function that registers a user in a database
func (profile *Profile) Register(ctx context.Context, req *profilepb.RegisterRequest) (*profilepb.RegisterResponse, error) {

	id, err := profile.strg.Lookup(ctx, req.UUID)
	if err != nil && err != profile.strg.ErrUserDoesNotExist() {
		return nil, err
	}

	if id != "" {
		return nil, ErrEmailInUse
	}

	err = profile.strg.Register(ctx, req.UUID, req.Username, req.FirstName, req.LastName, profile.defaultAvatar, profile.defaultBanner, "THIS IS A PLACEHOLDER", false, true)
	if err != nil {
		return nil, err
	}

	return &profilepb.RegisterResponse{UID: uuid, Token: token}, nil

}
