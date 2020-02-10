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

//ProfileToggleProjectVisible toggle's a user's available status
func (s *Server) ProfileToggleProjectVisible(ctx context.Context, req *backendpb.ProfileToggleProjectVisibleRequest) (*backendpb.ProfileToggleProjectVisibleResponse, error) {

	res, err := (*s.profileClient).ToggleProjectDisplay(ctx, &profilepb.ToggleDisplayRequest{
		UUID:      req.UUID,
		ProjectID: req.ProjectID,
	})
	if err != nil {
		return nil, err
	}
	if res == nil {
		logger.Println("this shouldnt happen")
		return nil, errNilResponse
	}

	//TODO
	//we may need to pull some data from project etc.

	return &backendpb.ProfileToggleProjectVisibleResponse{NewStatus: res.NewValue}, nil
}



//ProfileToggleProjectVisibleHTTP is an http wrapper for the signup function
func (s *Server) ProfileToggleProjectVisibleHTTP(w http.ResponseWriter, r *http.Request) {

	if !setCORS(&w, r) {
		return
	}

	uid, err := s.sesh.UID(r)
	if err != nil {
		logger.Println("Session error | err: ", err)
		data := packageResponse(nil, err)
		fmt.Fprintln(w, string(data))
		return
	}
	req := &backendpb.ProfileToggleProjectVisibleRequest{UUID: uid}

	if strings.Contains(r.Header.Get("Content-type"), "application/json") {
		r.Header.Set("Content-Type", "application/json")
		json.NewDecoder(r.Body).Decode(req)
	} else {
		r.ParseForm()
		req.ProjectID = r.Form.Get("tags")
	}

	resp, err := s.ProfileToggleProjectVisible(r.Context(), req)
	if err != nil {
		logger.Println("UUID: "+uid+" | err: ", err)
	}

	data := packageResponse(resp, err)

	fmt.Fprintln(w, string(data))

}
