package project

import (
	"context"

	"github.com/sea350/ustart_micro/backend/project/projectpb"
)

//GetApplicants pulls a list of all join requests
func (project *Project) GetApplicants(ctx context.Context, req *projectpb.GetApplicantsRequest) (*projectpb.GetApplicantsResponse, error) {

	reqsts, err := project.auth.GetRequests(ctx, req.ProjectID, req.RequesterID)
	if err != nil {
		return nil, err
	}

	return &projectpb.GetApplicantsResponse{ApplicantIDs: reqsts}, nil

}
