package profile

import (
	"context"

	"github.com/sea350/ustart_micro/backend/profile/profilepb"
)

//UpdateBasic updates a profile's basic information
func (profile *Profile) UpdateBasic(ctx context.Context, req *profilepb.UpdateBasicRequest) (*profilepb.UpdateBasicResponse, error) {

	err := profile.strg.UpdateBasicInfo(ctx, req.UUID, req.Firstname, req.Lastname, req.Description, req.Organization, req.DOB)
	//will still return successful if error
	//Be careful, it will only return one err even if multi step failure
	return &profilepb.UpdateBasicResponse{}, err

}
