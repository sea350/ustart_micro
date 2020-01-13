package auth

import (
	"context"
	"regexp"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/sea350/ustart_micro/backend/auth/authpb"
)

// Register is a generic register function that registers a user in a database
func (auth *Auth) Register(ctx context.Context, req *authpb.RegisterRequest) (*authpb.RegisterResponse, error) {

	id, err := auth.strg.Lookup(ctx, req.Email)
	if err != nil && err != auth.strg.ErrUserDoesNotExist() {
		return nil, err
	}

	if id != "" {
		return nil, ErrEmailInUse
	}

	rxEmail := regexp.MustCompile(`^[a-zA-Z0-9._-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,4}$`) //checks for the format of the email
	if !rxEmail.MatchString(req.Email) {
		return nil, ErrInvalidEmail
	}

	// encrypt the password
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), 12)
	if err != nil {
		return nil, err
	}

	//Generate uuid
	length := 16
	uuid := randString(length)
	inUse, err := auth.strg.IDLookup(ctx, uuid)
	if err != nil && err != auth.strg.ErrUserDoesNotExist() {
		return nil, err
	}
	for inUse {
		length++
		if length > 32 { //a uuid cannot be longer than 32 characters
			length = 32
		}
		uuid = randString(length)
		inUse, err = auth.strg.IDLookup(ctx, uuid)
		if err != nil && err != auth.strg.ErrUserDoesNotExist() {
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

	err = auth.strg.Register(ctx, uuid, req.Email, string(hashedPass), token, "user", time.Now().Format(auth.timeFormat), time.Now().Add(expireIn).Format(auth.timeFormat))
	if err != nil {
		return nil, err
	}

	return &authpb.RegisterResponse{UID: uuid, Token: token}, nil

}
