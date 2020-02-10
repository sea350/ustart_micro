package project

import (
	"context"

	"github.com/sea350/ustart_micro/backend/project/projectpb"
)

//RemoveMember removes a member
func (project *Project) RemoveMember(ctx context.Context, req *projectpb.RemoveMemberRequest) (*projectpb.RemoveMemberResponse, error) {

	err := project.auth.RemoveMember(ctx, req.ProjectID, req.RemoverID, req.RemovedID)
	if err != nil {
		return nil, err
	}

	return &projectpb.RemoveMemberResponse{}, nil
}
