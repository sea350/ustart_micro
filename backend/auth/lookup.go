package auth

import (
	"context"

	"github.com/sea350/ustart_micro/backend/auth/authpb"
)

//Lookup checks elastic search if a document exists with the criteria of the req
func (auth *Auth) Lookup(ctx context.Context, req *authpb.LookupRequest) (*authpb.LookupResponse, error) {
	id, err := auth.strg.Lookup(ctx, req.Email)

	var exists bool
	if id != "" {
		exists = true
	}

	return &authpb.LookupResponse{Exists: exists}, err
}
