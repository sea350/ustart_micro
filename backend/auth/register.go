package auth

import (
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"

	"github.com/sea350/ustart_mono/backend/auth/authpb"
)

// Register does what it does, works with pb req
func (eclient *Eclient) Register(ctx context.Context, req *authpb.RegisterRequest) (*authpb.RegisterResponse, error) {

	//Lock just to make sure no two people can sign up with the same email at the same time
	newUserLock.Lock()
	defer newUserLock.Unlock()

	// make sure email is not in use
	exists := eclient.EmailExists(ctx, req.Email)
	if exists {
		return &authpb.RegisterResponse{}, errors.New("Email in use")
	}

	//before instering into database make sure the index exists
	eclient.IndexExists(ctx)

	// encrypt the password
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), 12)
	if err != nil {
		return &authpb.RegisterResponse{}, err
	}

	store := authpb.Stored{Email: req.Email, Password: string(hashedPass)}

	//GENERATE TOKEN HERE

	//NOT SURE ABOUT THE BODY JSON
	newUser, err := eclient.client.Index().
		Index(eIndex).
		Type(eType).
		BodyJson(store).
		Do(ctx)

	if err != nil {
		return &authpb.RegisterResponse{}, err
	}

	return &authpb.RegisterResponse{UID: newUser.Id}, err

}
