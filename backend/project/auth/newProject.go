package auth

import (
	"context"
	"time"
)

//NewProject sets fundamental data structures necessary for a new project
func (auth *Auth) NewProject(ctx context.Context, projectID, creatorID string) error {
	//register the creator
	err := auth.Strg.AddMember(ctx, creatorID, projectID, "Creator", time.Now().Format(auth.timeFormat))
	if err != nil {
		return err
	}

	//add the default roles member and moderator
	err = auth.Strg.AddRole(ctx, "Member", projectID, false, false, false, false, false, false)
	if err != nil {
		return err
	}

	err = auth.Strg.AddRole(ctx, "Moderator", projectID, true, true, true, true, true, true)

	return err
}
