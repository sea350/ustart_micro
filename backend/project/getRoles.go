package project

import (
	"context"

	"github.com/sea350/ustart_micro/backend/project/projectpb"
)

//GetRoles pulls a list of all join requests
func (project *Project) GetRoles(ctx context.Context, req *projectpb.GetRolesRequest) (*projectpb.GetRolesResponse, error) {

	reqsts, err := project.auth.GetRoles(ctx, req.ProjectID, req.UserID)
	if err != nil {
		return nil, err
	}

	return &projectpb.GetRolesResponse{Roles: reqsts}, nil
}
