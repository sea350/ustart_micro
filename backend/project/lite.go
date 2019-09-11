package project

import (
	"context"

	"github.com/sea350/ustart_micro/backend/project/projectpb"
)

// Lite retreives minimal profile data based off of a uuid
func (project *Project) Lite(ctx context.Context, req *projectpb.LiteRequest) (*projectpb.LiteResponse, error) {

	proj, err := project.strg.Lookup(ctx, req.PID)
	if err != nil {
		return nil, err
	}

	return &projectpb.LiteResponse{
		PID:         proj.PID,
		CustomURL:   proj.CustomURL,
		Name:        proj.Name,
		Description: proj.Description,
		Avatar:      proj.Avatar,
	}, nil

}
