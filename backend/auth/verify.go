package auth

import (
	"context"
	"time"

	"github.com/sea350/ustart_micro/backend/auth/authpb"
)

//Verify checks a user's token and verifies them if they are a match
func (auth *Auth) Verify(ctx context.Context, req *authpb.VerifyRequest) (*authpb.VerifyResponse, error) {
	token, expiration, err := auth.strg.GetToken(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	expTime, err := time.Parse(auth.timeFormat, expiration)
	if err != nil {
		return &authpb.VerifyResponse{}, err
	}

	if token == req.Token && time.Now().Before(expTime) {
		err := auth.strg.Validate(ctx, req.Email, true)
		if err != nil {
			return nil, err
		}

		return &authpb.VerifyResponse{}, auth.strg.SetToken(ctx, req.Email, "", time.Time{}.Format(auth.timeFormat))
	}

	return &authpb.VerifyResponse{}, ErrInvalidToken
}
