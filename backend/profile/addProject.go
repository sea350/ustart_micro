package profile

import (
	"context"

	"github.com/sea350/ustart_micro/backend/profile/profilepb"
)

//AddProject add a new project to the display (default = visibile)
func (profile *Profile) AddProject(ctx context.Context, req *profilepb.AddProjectRequest) (*profilepb.AddProjectResponse, error) {
	modProjectLock.Lock()
	defer modProjectLock.Unlock()

	prof, err := profile.strg.Lookup(ctx, req.UUID)
	if err != nil {
		return nil, err
	}

	prof.Projects = append(prof.Projects, &profilepb.ProjectDisplay{
		ProjectID: req.ProjectID,
		Visible:   true,
	})

	err = profile.strg.UpdateProjects(ctx, req.UUID, prof.Projects)
	if err != nil {
		return nil, err
	}

	return &profilepb.AddProjectResponse{}, nil

}
