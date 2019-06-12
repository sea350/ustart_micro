package backend

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

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
		log.Print("this shouldnt happen")
		return nil, err
	}

	//TODO:
	//Education history

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

//EXPERIMENTAL

// SignupHTTP is an http wrapper for the signup function
func (s *Server) SignupHTTP(w http.ResponseWriter, r *http.Request) {

	regCtx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

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
	}

	ret := make(map[string]interface{})

	resp, err := s.Signup(regCtx, req)
	if resp != nil {
		ret["response"] = resp
	} else {
		ret["response"] = ""
	}
	if err != nil {
		ret["error"] = err.Error()
	} else {
		ret["error"] = ""
	}

	data, err := json.Marshal(ret)
	if err != nil {
		log.Println("Problem martialing return data", err)
	}

	fmt.Fprintln(w, string(data))

}
