package auth

import (
	"context"

	"github.com/sea350/ustart_mono/backend/auth/authpb"
)

//Lookup checks elastic search if a document exists with the criteria of the req
func (auth *Auth) Lookup(ctx context.Context, req *authpb.LookupRequest) (*authpb.LookupResponse, error) {
	exists, err := auth.strg.Lookup(ctx, req.Email)

	return &authpb.LookupResponse{Exists: exists}, err
}
