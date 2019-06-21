package backend

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/sea350/ustart_micro/backend/backendpb"
	"github.com/sea350/ustart_micro/backend/profile/profilepb"
)

//PublicProfile retrieves all publicly accessible profile data
func (s *Server) PublicProfile(ctx context.Context, req *backendpb.ProfileRequest) (*backendpb.ProfileResponse, error) {

	pullRes, err := (*s.profileClient).Pull(ctx, &profilepb.PullRequest{
		Username: req.Username,
	})
	if err != nil {
		return nil, err
	}
	if pullRes == nil {
		logger.Println("this shouldnt happen")
		return nil, err
	}

	//TODO
	//we may need to pull some data from project etc.

	return &backendpb.ProfileResponse{UserData: convertProfileUserToBackendUser(*pullRes.UserProfile)}, nil
}

//EXPERIMENTAL

//PublicProfileHTTP is an http wrapper for the signup function
func (s *Server) PublicProfileHTTP(w http.ResponseWriter, r *http.Request) {

	if !setCORS(&w, r) {
		return
	}

	req := &backendpb.ProfileRequest{}

	if strings.Contains(r.Header.Get("Content-type"), "application/json") {
		r.Header.Set("Content-Type", "application/json")
		json.NewDecoder(r.Body).Decode(req)
	} else {
		r.ParseForm()
		req.Username = r.Form.Get("username")
	}

	ret := make(map[string]interface{})

	resp, err := s.PublicProfile(r.Context(), req)
	if resp != nil {
		ret["response"] = resp
	} else {
		ret["response"] = ""
	}
	if err != nil {
		ret["error"] = err.Error()
		logger.Println("Usernmae: "+req.Username+" | err: ", err)
	} else {
		ret["error"] = ""
	}

	data, err := json.Marshal(ret)
	if err != nil {
		logger.Println("Problem martialing return data", err)
	}

	fmt.Fprintln(w, string(data))

}
