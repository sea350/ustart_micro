package auth

import (
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"

	"github.com/sea350/ustart_mono/backend/auth/authpb"
)

// Register does what it does, works with pb req
func (eauth *ElasticAuth) Register(ctx context.Context, req *authpb.RegisterRequest) (*authpb.RegisterResponse, error) {

	//Lock just to make sure no two people can sign up with the same email at the same time
	newUserLock.Lock()
	defer newUserLock.Unlock()

	// make sure email is not in use
	res, err := eauth.Lookup(ctx, &authpb.LookupRequest{Email: req.Email})
	if err != nil {
		return &authpb.RegisterResponse{}, err
	}
	if res.Exists {
		return &authpb.RegisterResponse{}, errors.New("Email in use")
	}

	//before instering into database make sure the index exists
	exists, err := eauth.client.IndexExists(eauth.eIndex).Do(ctx)
	if err != nil {
		panic(err)
	}

	// If the index doesn't exist, create it. there shouldnt be any errors but if there are they are critical
	if !exists {
		createIndex, err := eauth.client.CreateIndex(eauth.eIndex).BodyString("").Do(ctx)
		if err != nil {
			panic(err)
		}
		// TODO fix this.
		if !createIndex.Acknowledged {
			panic(err)
		}
	}

	// encrypt the password
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), 12)
	if err != nil {
		return &authpb.RegisterResponse{}, err
	}

	store := authpb.Stored{Email: req.Email, Password: string(hashedPass)}

	//GENERATE TOKEN HERE

	//NOT SURE ABOUT THE BODY JSON
	newUser, err := eauth.client.Index().
		Index(eauth.eIndex).
		Type(eauth.eType).
		BodyJson(store).
		Do(ctx)

	if err != nil {
		return &authpb.RegisterResponse{}, err
	}

	return &authpb.RegisterResponse{UID: newUser.Id}, err

}
