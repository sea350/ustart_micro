package profile

import (
	"context"

	"github.com/sea350/ustart_micro/backend/profile/profilepb"
)

//UpdateBanner updates a profile's banner
func (profile *Profile) UpdateBanner(ctx context.Context, req *profilepb.UpdateBannerRequest) (*profilepb.UpdateBannerResponse, error) {

	err := profile.strg.UpdateBanner(ctx, req.UUID, req.NewBannerLink)
	if err != nil {
		return nil, err
	}
	return &profilepb.UpdateBannerResponse{}, nil

}
