package auth

import (
	"context"
	"time"

	"github.com/sea350/ustart_micro/backend/auth/authpb"
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

	expTime, err := time.Parse(auth.timeFormat, expiration)
	if err != nil {
		return &authpb.RecoverResponse{}, err
	}

	if token == req.Token && time.Now().Before(expTime) {
		err := auth.strg.ChangePassword(ctx, req.Email, string(hashedPass))
		if err != nil {
			return nil, err
		}
		//setting expiration time
		var expireIn time.Duration
		if auth.tokenExpiration == 0 {
			expireIn = 48 * time.Hour //default expire time is 2 days
		} else {
			expireIn = time.Duration(auth.tokenExpiration) * time.Hour
		}
		return &authpb.RecoverResponse{}, auth.strg.SetToken(ctx, req.Email, "", time.Now().Add(expireIn).Format(auth.timeFormat))
	}

	return &authpb.RecoverResponse{}, ErrInvalidToken
}
