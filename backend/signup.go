package backend

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/sea350/ustart_micro/backend/backendpb"
)

//Signup registers a new user
func (s *Server) Signup(ctx context.Context, req *backendpb.SignupRequest) (*backendpb.SignupResponse, error) {

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
