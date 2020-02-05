package backend

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"strconv"
	"strings"

	"github.com/sea350/ustart_micro/backend/auth/authpb"
	"github.com/sea350/ustart_micro/backend/backendpb"
)

//Login verifies a user's credentials and keeps track of their session
func (s *Server) Login(ctx context.Context, req *backendpb.LoginRequest) (*backendpb.LoginResponse, error) {

	//Authenticate credentials
	resAuth, err := (*s.authClient).Authenticate(ctx, &authpb.AuthenticateRequest{
		Email:     req.Identifier,
		Challenge: req.Challenge,
	})
	if err != nil {
		return nil, err
	}
	if resAuth == nil {
		logger.Println("this shouldnt happen")
		return nil, errNilResponse
	}
	//Credentials verified

	//Session activated in http wrapper of this func

	return &backendpb.LoginResponse{UUID: resAuth.UUID}, nil
}

//EXPERIMENTAL

//LoginHTTP is an http wrapper for the login function, also initializes a new http session
func (s *Server) LoginHTTP(w http.ResponseWriter, r *http.Request) {

	if !setCORS(&w, r) {
		return
	}

	req := &backendpb.LoginRequest{}
	var rememberMe bool

	if strings.Contains(r.Header.Get("Content-type"), "application/json") {
		r.Header.Set("Content-Type", "application/json")
		json.NewDecoder(r.Body).Decode(req)
	} else {
		r.ParseForm()
		req.Identifier = r.Form.Get("email")
		req.Challenge = r.Form.Get("password")
		rememberMe, _ = strconv.ParseBool(r.Form.Get("rememberMe"))
	}

	resp, err := s.Login(r.Context(), req)

	var sessID string
	if resp != nil && err == nil {
		ip, _, _ := net.SplitHostPort(r.RemoteAddr)
		sessID, err = s.sesh.Start(resp.UUID, "", ip, rememberMe, &w, r)
	}
	resp.SessionID = sessID
	if err != nil {
		logger.Println("Login identifier: "+req.Identifier+" | err: ", err)
	}

	data := packageResponse(resp, err)

	fmt.Fprintln(w, string(data))

}
