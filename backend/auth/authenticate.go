package auth

import (
	"context"

	"github.com/sea350/ustart_micro/backend/auth/authpb"
	"golang.org/x/crypto/bcrypt"
)

//Authenticate validates an authentication request, works with pb req
func (auth *Auth) Authenticate(ctx context.Context, req *authpb.AuthenticateRequest) (*authpb.AuthenticateResponse, error) {

	pass, err := auth.strg.GetPassword(ctx, req.Email)
	if err != nil {
		return &authpb.AuthenticateResponse{}, err
	}

	//validate the password
	err = bcrypt.CompareHashAndPassword([]byte(pass), []byte(req.Challenge))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		return &authpb.AuthenticateResponse{}, ErrIncorrectPassword
	}
	if err != nil {
		return &authpb.AuthenticateResponse{}, err
	}

	id, err := auth.strg.Lookup(ctx, req.Email)
	if err != nil {
		return &authpb.AuthenticateResponse{}, err
	}

	return &authpb.AuthenticateResponse{UUID: id}, nil
}
