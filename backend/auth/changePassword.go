package auth

import (
	"context"

	"golang.org/x/crypto/bcrypt"

	"github.com/sea350/ustart_micro/backend/auth/authpb"
)

//ChangePassword automatically authenticates and then updates
func (auth *Auth) ChangePassword(ctx context.Context, req *authpb.ChangePasswordRequest) (*authpb.ChangePasswordResponse, error) {
	//submit an authentication request to confirm valid acc info
	_, err := auth.Authenticate(ctx, &authpb.AuthenticateRequest{Email: req.Email, Challenge: req.Challenge})
	if err != nil {
		return nil, err
	}
	// encrypt the password
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), 12)
	if err != nil {
		return nil, err
	}

	err = auth.strg.ChangePassword(ctx, req.Email, string(hashedPass))
	if err != nil {
		return nil, err
	}

	return &authpb.ChangePasswordResponse{}, nil

}
