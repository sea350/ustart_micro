package project

import (
	"context"

	"github.com/sea350/ustart_micro/backend/project/projectpb"
)

// Pull retreives all project data based off of a project url
func (project *Project) Pull(ctx context.Context, req *projectpb.PullRequest) (*projectpb.PullResponse, error) {

	proj, err := project.strg.LookupCustomURL(ctx, req.CustomURL)
	if err != nil {
		return nil, err
	}

	return &projectpb.PullResponse{ProjectPage: &proj}, nil

}
