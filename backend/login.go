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

//Login registers a new user
func (s *Server) Login(ctx context.Context, req *backendpb.LoginRequest) (*backendpb.LoginResponse, error) {
	//TODO:
	//all of it

	return &backendpb.LoginResponse{}, nil
}

//EXPERIMENTAL

//LoginHTTP is an http wrapper for the signup function
func (s *Server) LoginHTTP(w http.ResponseWriter, r *http.Request) {

	regCtx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	if !setCORS(&w, r) {
		return
	}

	req := &backendpb.LoginRequest{}

	if strings.Contains(r.Header.Get("Content-type"), "application/json") {
		r.Header.Set("Content-Type", "application/json")
		json.NewDecoder(r.Body).Decode(req)
	} else {
		r.ParseForm()
		req.Identifier = r.Form.Get("email")
		req.Challenge = r.Form.Get("password")
	}

	ret := make(map[string]interface{})

	resp, err := s.Login(regCtx, req)
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
