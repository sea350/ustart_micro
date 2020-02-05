package profile

import (
	"context"

	"github.com/sea350/ustart_micro/backend/profile/profilepb"
)

//ToggleProjectDisplay turns a project's visibility on/off
func (profile *Profile) ToggleProjectDisplay(ctx context.Context, req *profilepb.ToggleDisplayRequest) (*profilepb.ToggleDisplayResponse, error) {
	modProjectLock.Lock()
	defer modProjectLock.Unlock()

	prof, err := profile.strg.Lookup(ctx, req.UUID)
	if err != nil {
		return nil, err
	}

	var broken bool
	var newStatus bool
	for _, proj := range prof.Projects {
		if proj.ProjectID == req.ProjectID {
			proj.Visible = !proj.Visible
			newStatus = proj.Visible
			broken = true
			break
		}
	}

	if !broken {
		return nil, ErrNoProjectFound
	}

	err = profile.strg.UpdateProjects(ctx, req.UUID, prof.Projects)
	if err != nil {
		return nil, err
	}

	return &profilepb.ToggleDisplayResponse{NewValue: newStatus}, nil

}
