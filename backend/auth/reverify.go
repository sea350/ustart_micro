package auth

import (
	"context"
	"time"

	"github.com/sea350/ustart_mono/backend/auth/authpb"
)

//Reverify resets a user's token and expiration date
func (auth *Auth) Reverify(ctx context.Context, req *authpb.ReverifyRequest) (*authpb.ReverifyResponse, error) {

	//setting expiration time
	var expireIn time.Duration
	if auth.tokenExpiration == 0 {
		expireIn = 48 * time.Hour //default expire time is 2 days
	} else {
		expireIn = time.Duration(auth.tokenExpiration) * time.Hour
	}

	token := randString(16) //generate new token

	err := auth.strg.SetToken(ctx, req.Email, token, time.Now().Add(expireIn))
	if err != nil {
		return &authpb.ReverifyResponse{}, err
	}

	return &authpb.ReverifyResponse{Token: token}, nil
}
