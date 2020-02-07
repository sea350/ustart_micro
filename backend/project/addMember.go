package project

import (
	"context"

	"github.com/sea350/ustart_micro/backend/project/projectpb"
)

//AddMember adds a new member with all the requisite checks
func (project *Project) AddMember(ctx context.Context, req *projectpb.AddMemberRequest) (*projectpb.AddMemberResponse, error) {

	err := project.auth.AddMember(ctx, req.ProjectID, req.AdderID, req.NewMemberID, req.RoleName)
	if err != nil {
		return nil, err
	}

	return &projectpb.AddMemberResponse{}, nil
}
