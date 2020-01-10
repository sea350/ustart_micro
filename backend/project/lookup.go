package project

import (
	"context"

	"github.com/sea350/ustart_micro/backend/project/projectpb"
)

// Lookup checks to see if a project url is in use
func (project *Project) Lookup(ctx context.Context, req *projectpb.LookupRequest) (*projectpb.LookupResponse, error) {

	_, err := project.strg.LookupCustomURL(ctx, req.CustomURL)
	if err == project.strg.ErrProjectDoesNotExist() {
		return &projectpb.LookupResponse{Exists: false}, nil
	}
	if err != nil {
		return nil, err
	}

	return &projectpb.LookupResponse{Exists: true}, nil

}
