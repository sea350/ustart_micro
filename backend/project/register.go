package project

import (
	"context"

	"github.com/sea350/ustart_micro/backend/project/projectpb"
)

// Register is a generic register function that registers a user in a database
func (project *Project) Register(ctx context.Context, req *projectpb.RegisterRequest) (*projectpb.RegisterResponse, error) {

	//Why would you even need to use this, YOU CANT HAVE AN ID OF AN OBJECT UR TRYING TO MAKE
	//MIN REEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEE
	// _, err := project.strg.Lookup(ctx, req.PID)
	// if err != nil && err != project.strg.ErrProjectDoesNotExist() {
	// 	return nil, err
	// }
	// if err == nil {
	// 	return nil, ErrProjectExists
	// }

	_, err := project.strg.LookupCustomURL(ctx, req.CustomURL)
	if err != nil && err != project.strg.ErrProjectDoesNotExist() {
		return nil, err
	}
	if err == nil {
		return nil, ErrCustomURLInUse
	}

	pid, err := project.strg.Register(ctx, req.CustomURL, req.Name, "category", project.defaultAvatar, project.defaultBanner, "creationdate", req.School, []string{}, []string{}, []string{})
	if err != nil {
		return nil, err
	}

	err = project.auth.NewProject(ctx, pid, req.CreatorID)

	return &projectpb.RegisterResponse{ProjectID: pid}, err

}
