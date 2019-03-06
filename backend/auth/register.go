package auth

import (
	"context"

	"golang.org/x/crypto/bcrypt"

	"github.com/sea350/ustart_mono/backend/auth/authpb"
)

// Register does what it does
func (auth *Auth) Register(ctx context.Context, req *authpb.RegisterRequest) (*authpb.RegisterResponse, error) {

	//todo:
	//CHECK IF SOMEONE IS ALREADY REGISTERED USNIG GIVEN CREDENTIALS

	// encrypt the password
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), 12)
	if err != nil {
		return nil, err
	}

	//todo:
	//DONT FORGET ABOUT SETTING UP UUID
	//INSERT ACC TYEP SWITCH HERE

	err = auth.strg.Register(ctx, req.Email, string(hashedPass), "INSERT TOKEN HERE", "user")

	return &authpb.RegisterResponse{}, err

}
