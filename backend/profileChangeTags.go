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

//ProfileChangeTags toggle's a user's available status
func (s *Server) ProfileChangeTags(ctx context.Context, req *backendpb.ProfileChangeTagsRequest) (*backendpb.ProfileChangeTagsResponse, error) {

	res, err := (*s.profileClient).UpdateTags(ctx, &profilepb.UpdateTagsRequest{
		UUID: req.UUID,
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

	return &backendpb.ProfileChangeTagsResponse{}, nil
}

//EXPERIMENTAL

//ProfileChangeTagsHTTP is an http wrapper for the signup function
func (s *Server) ProfileChangeTagsHTTP(w http.ResponseWriter, r *http.Request) {

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
	req := &backendpb.ProfileChangeTagsRequest{UUID: uid}

	if strings.Contains(r.Header.Get("Content-type"), "application/json") {
		r.Header.Set("Content-Type", "application/json")
		json.NewDecoder(r.Body).Decode(req)
	} else {
		r.ParseForm()
		err = json.Unmarshal([]byte(r.Form.Get("tags")), &req.Tags)
		if err != nil {
			logger.Println("Form unmarshall error, uuid: "+uid+" | err: ", err)
			data := packageResponse(nil, err)
			fmt.Fprintln(w, string(data))
			return
		}
	}

	resp, err := s.ProfileChangeTags(r.Context(), req)
	if err != nil {
		logger.Println("UUID: "+uid+" | err: ", err)
	}

	data := packageResponse(resp, err)

	fmt.Fprintln(w, string(data))

}
