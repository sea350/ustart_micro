package profile

import (
	"context"

	"github.com/sea350/ustart_micro/backend/profile/profilepb"
)

//RemoveProject eliminate a project from the display
func (profile *Profile) RemoveProject(ctx context.Context, req *profilepb.RemoveProjectRequest) (*profilepb.RemoveProjectResponse, error) {
	modProjectLock.Lock()
	defer modProjectLock.Unlock()

	prof, err := profile.strg.Lookup(ctx, req.UUID)
	if err != nil {
		return nil, err
	}

	var broken bool
	for i, proj := range prof.Projects {
		if proj.ProjectID == req.ProjectID {
			prof.Projects = append(prof.Projects[:i], prof.Projects[i+1:]...)
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

	return &profilepb.RemoveProjectResponse{}, nil

}
