package auth

import (
	"context"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/sea350/ustart_mono/backend/auth/authpb"
	"github.com/sea350/ustart_mono/backend/auth/storage/sql"
)

// Register does what it does
func (auth *Auth) Register(ctx context.Context, req *authpb.RegisterRequest) (*authpb.RegisterResponse, error) {

	//todo:
	//CHECK IF SOMEONE IS ALREADY REGISTERED USNIG GIVEN CREDENTIALS
	//VALIDATE THE EMAIL ACCORDING TO THE REGISTRATION FUNCTION

	// encrypt the password
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), 12)
	if err != nil {
		return nil, err
	}

	//Generate uuid
	length := 16
	uuid := randString(length)
	inUse, err := auth.strg.IDLookup(ctx, uuid)
	if err != nil && err != sqlstore.ErrUserDoesNotExist { //todo: MAKE THIS DATABASE AGNOSTIC
		return nil, err
	}
	for inUse {
		length++
		if length > 32 { //a uuid cannot be longer than 32 characters
			length = 32
		}
		uuid = randString(length)
		inUse, err = auth.strg.IDLookup(ctx, uuid)
		if err != nil && err != sqlstore.ErrUserDoesNotExist { //todo: MAKE THIS DATABASE AGNOSTIC
			return nil, err
		}
	}

	//generating token
	token := randString(16)

	//setting expiration time
	var expireIn time.Duration
	if auth.tokenExpiration == 0 {
		expireIn = 48 * time.Hour //default expire time is 2 days
	} else {
		expireIn = time.Duration(auth.tokenExpiration) * time.Hour
	}

	err = auth.strg.Register(ctx, uuid, req.Email, string(hashedPass), token, "user", time.Now().Add(expireIn))

	return &authpb.RegisterResponse{UID: uuid, Token: token}, err

}
