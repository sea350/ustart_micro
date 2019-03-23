package auth

import (
	"context"
	"time"

	"github.com/sea350/ustart_mono/backend/auth/authpb"
	"golang.org/x/crypto/bcrypt"
)

//RecoverPassword checks a user's token and updates with a new password
func (auth *Auth) RecoverPassword(ctx context.Context, req *authpb.RecoverRequest) (*authpb.RecoverResponse, error) {
	token, expiration, err := auth.strg.GetToken(ctx, req.Email)
	if err != nil {
		return &authpb.RecoverResponse{}, err
	}

	// encrypt the password
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), 12)
	if err != nil {
		return nil, err
	}

	if token == req.Token && time.Now().Before(expiration) {
		err := auth.strg.ChangePassword(ctx, req.Email, string(hashedPass))
		if err != nil {
			return nil, err
		}

		return &authpb.RecoverResponse{}, auth.strg.SetToken(ctx, req.Email, "", time.Time{})
	}

	return &authpb.RecoverResponse{}, ErrInvalidToken
}
