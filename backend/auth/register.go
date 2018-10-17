package auth

import (
	"context"
	"errors"
	"strings"

	"golang.org/x/crypto/bcrypt"

	"github.com/olivere/elastic"
	"github.com/sea350/ustart_mono/backend/auth/authpb"
)

// Register does what it does
func (eclient *Eclient) Register(ctx context.Context, req *authpb.RegisterRequest) (*authpb.RegisterResponse, error) {
	// make sure email is not in use
	termQuery := elastic.NewTermQuery("Email", strings.ToLower(req.Email))
	res, err := eclient.Search().
		Index(eIndex).
		Query(termQuery).
		Do(ctx)

	if err != nil {
		return &authpb.RegisterResponse{}, err
	}

	if res.Hits.TotalHits > 0 {
		return &authpb.RegisterResponse{}, errors.New("Email is in use")
	}

	exists, err := eclient.IndexExists(eIndex).Do(ctx)
	if err != nil {
		return ID, err
	}

	// If the index doesn't exist, create it and return error.
	if !exists {
		createIndex, err := eclient.CreateIndex(eIndex).BodyString("").Do(ctx)
		if err != nil {
			panic(err)
		}
		// TODO fix this.
		if !createIndex.Acknowledged {
			return &authpb.RegisterResponse{}, errors.New("Index not acknowledged")
		}
	}

	// encrypt the password
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), 12)
	if err != nil {
		return &authpb.RegisterResponse{}, err
	}

	//NOT SURE ABOUT THE BODY JSON
	newUser, err := eclient.Index().
		Index(eIndex).
		Type(eType).
		BodyJson(req).
		Do(ctx)

	if err != nil {
		return &authpb.RegisterResponse{}, err
	}

	return &authpb.RegisterResponse{UID: newUser.Id}, err

}
