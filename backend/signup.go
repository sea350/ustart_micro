package backend

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/sea350/ustart_micro/backend/profile/profilepb"

	"github.com/sea350/ustart_micro/backend/auth/authpb"
	"github.com/sea350/ustart_micro/backend/backendpb"
)

//Signup registers a new user
func (s *Server) Signup(ctx context.Context, req *backendpb.SignupRequest) (*backendpb.SignupResponse, error) {

	resAuth, err := (*s.authClient).Register(ctx, &authpb.RegisterRequest{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}
	if resAuth == nil {
		logger.Println("this shouldnt happen")
		return nil, errNilResponse
	}

	//TODO:
	//Education history may need to be expanded on

	_, err = (*s.profileClient).Register(ctx, &profilepb.RegisterRequest{
		UUID:      resAuth.UID,
		Username:  req.Username,
		FirstName: req.FirstName,
		LastName:  req.LastName,
	})
	if err != nil {
		return nil, err
	}

	//TODO:
	//send token to emailer
	//register with any other subservices

	return &backendpb.SignupResponse{}, nil
}



// SignupHTTP is an http wrapper for the signup function
func (s *Server) SignupHTTP(w http.ResponseWriter, r *http.Request) {

	if !setCORS(&w, r) {
		return
	}

	req := &backendpb.SignupRequest{}

	if strings.Contains(r.Header.Get("Content-type"), "application/json") {
		r.Header.Set("Content-Type", "application/json")
		json.NewDecoder(r.Body).Decode(req)
	} else {
		r.ParseForm()
		req.Email = r.Form.Get("email")
		req.Password = r.Form.Get("password")
		req.Username = r.Form.Get("username")
		req.FirstName = r.Form.Get("firstname")
		req.LastName = r.Form.Get("lastname")
	}

	resp, err := s.Signup(r.Context(), req)
	if err != nil {
		logger.Println("Email: "+req.Email+" | err: ", err)
	}

	data := packageResponse(resp, err)

	fmt.Fprintln(w, string(data))

}
