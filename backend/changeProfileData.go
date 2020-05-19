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

//ChangeProfileData updates a single profile's basic information
func (s *Server) ChangeProfileData(ctx context.Context, req *backendpb.ChangeBasicDataRequest) (*backendpb.ChangeBasicDataResponse, error) {

	pullRes, err := (*s.profileClient).UpdateBasic(ctx, &profilepb.UpdateBasicRequest{
		UUID:         req.UUID,
		Firstname:    req.Firstname,
		Lastname:     req.Lastname,
		Description:  req.Description,
		Organization: req.Organization,
		DOB:          req.DOB,
	})
	if err != nil {
		return nil, err
	}
	if pullRes == nil {
		logger.Println("this shouldnt happen")
		return nil, errNilResponse
	}

	//TODO
	//we may need to pull some data from project etc.

	return &backendpb.ChangeBasicDataResponse{}, nil
}

//ChangeProfileDataHTTP is an http wrapper for the ChangeProfileData function
func (s *Server) ChangeProfileDataHTTP(w http.ResponseWriter, r *http.Request) {
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

	req := &backendpb.ChangeBasicDataRequest{}

	if strings.Contains(r.Header.Get("Content-type"), "application/json") {
		r.Header.Set("Content-Type", "application/json")
		json.NewDecoder(r.Body).Decode(req)
	} else {
		r.ParseForm()
		req.Firstname = r.Form.Get("firstname")
		req.Firstname = r.Form.Get("lastname")
		req.Firstname = r.Form.Get("description")
		req.Firstname = r.Form.Get("organization")
		req.Firstname = r.Form.Get("dob")
	}

	resp, err := s.ChangeProfileData(r.Context(), req)
	if err != nil {
		logger.Println("UUID: "+uid+" | err: ", err)
	}

	data := packageResponse(resp, err)

	fmt.Fprintln(w, string(data))

}
