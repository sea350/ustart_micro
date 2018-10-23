package auth

import (
	"context"

	"golang.org/x/crypto/bcrypt"

	"github.com/sea350/ustart_mono/backend/auth/authpb"
)

//ChangePassword automatically authenticates and then updates
func (eclient *ElasticAuth) ChangePassword(ctx context.Context, req *authpb.ChangePasswordRequest) (*authpb.ChangePasswordResponse, error) {
	//submit an authentication request to confirm valid acc info
	res, err := eclient.Authenticate(ctx, &authpb.AuthenticateRequest{Email: req.Email, Challenge: req.Challenge})
	if err != nil {
		return &authpb.ChangePasswordResponse{}, err
	}
	// encrypt the password
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), 12)
	if err != nil {
		return &authpb.ChangePasswordResponse{}, err
	}

	_, err = eclient.client.Update().
		Index(eIndex).
		Id(res.UID).
		Doc(map[string]interface{}{"Password": hashedPass}).
		Do(ctx)

	return &authpb.ChangePasswordResponse{}, err

}
