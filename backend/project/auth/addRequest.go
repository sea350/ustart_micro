package auth

import (
	"context"
	"time"
)

//AddRequest checks whether the user is currently part of the project and sends the request to the queue if not
func (auth *Auth) AddRequest(ctx context.Context, projectID, userID string) error {

	//first checks if user currently is part of the project
	role, err := auth.FindUserRole(ctx, projectID, userID)
	if err == nil {
		return errAlreadyExists //user is already a member
	}

	//Adds request with the time of request
	auth.Strg.AddRequest(ctx, projectID, userID, time.Now().Format(auth.timeFormat))
	if err != nil {
		return err
	}

	return nil
}
