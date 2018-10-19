package auth

import (
	"context"
	"encoding/json"
	"errors"
	"strings"

	elastic "github.com/olivere/elastic"
	"github.com/sea350/ustart_mono/backend/auth/authpb"
	"golang.org/x/crypto/bcrypt"
)

//Authenticate validates an authentication request, works with pb req
func (eclient *Eclient) Authenticate(ctx context.Context, req *authpb.AuthenticateRequest) (*authpb.AuthenticationResponce, error) {

	//pull soted data attached to the email
	query := elastic.NewTermQuery("Email", strings.ToLower(req.Email))
	res, err := eclient.client.Search().
		Index(eIndex).
		Query(query).
		Do(ctx)

	if err != nil {
		return &authpb.AuthenticationResponce{}, err
	}

	//there should never be more than one result. If there is, there is an issue
	if res.Hits.TotalHits > 1 {
		return &authpb.AuthenticationResponce{}, errors.New("More than one result")
	}

	var usr authpb.Stored
	var ID string //this is left here in case it is needed in the future

	for _, element := range res.Hits.Hits {
		ID = element.Id
		err := json.Unmarshal(*element.Source, &usr)
		if err != nil {
			return &authpb.AuthenticationResponce{}, err
		}
		break
	}

	//validate the password
	valid := bcrypt.CompareHashAndPassword(usr.Password, req.Challenge)
	if !valid {
		return &authpb.AuthenticationResponce{}, errors.New("Invalid password")
	}

	return &authpb.AuthenticationResponce{Success: true}, nil
}
