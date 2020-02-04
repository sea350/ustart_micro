package profile

import (
	"context"

	"github.com/sea350/ustart_micro/backend/profile/profilepb"
)

//UpdateIcon updates a profile's icon
func (profile *Profile) UpdateIcon(ctx context.Context, req *profilepb.UpdateIconRequest) (*profilepb.UpdateIconResponse, error) {

	err := profile.strg.UpdateAvatar(ctx, req.UUID, req.NewIconLink)
	if err != nil {
		return nil, err
	}
	return &profilepb.UpdateIconResponse{}, nil

}
