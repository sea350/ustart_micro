package project

import (
	"context"

	"github.com/sea350/ustart_micro/backend/project/projectpb"
)

//Reject adds a new member with all the requisite checks
func (project *Project) Reject(ctx context.Context, req *projectpb.RejectRequest) (*projectpb.RejectResponse, error) {

	err := project.auth.RemoveRequest(ctx, req.ProjectID, req.RemovedID, req.RemoverID)
	if err != nil {
		return nil, err
	}

	return &projectpb.RejectResponse{}, nil
}
