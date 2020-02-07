package backend

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/sea350/ustart_micro/backend/backendpb"
	"github.com/sea350/ustart_micro/backend/profile/profilepb"
	"github.com/sea350/ustart_micro/backend/project/projectpb"
)

//GetJoiners retrieves the lite profiles of those who have requested to join a project
func (s *Server) GetJoiners(ctx context.Context, req *backendpb.GetJoinersRequest) (*backendpb.GetJoinersResponse, error) {

	appRes, err := (*s.projectClient).GetApplicants(ctx, &projectpb.GetApplicantsRequest{
		ProjectID:   req.ProjectID,
		RequesterID: req.UUID,
	})
	if err != nil {
		return nil, err
	}
	if appRes == nil {
		logger.Println("this shouldnt happen")
		return nil, errNilResponse
	}

	var applicants []*backendpb.LiteProfile
	for _, appID := range appRes.ApplicantIDs {
		liteRes, errTemp := (*s.profileClient).Lite(ctx, &profilepb.LiteRequest{UUID: appID})
		if errTemp != nil {
			err = errProblemLoadingArr
		} else if liteRes == nil {
			logger.Println("this shouldnt happen, associated ID:", appID)
		} else {
			applicants = append(applicants, convertLiteResToLiteProfile(liteRes))
		}
	}

	return &backendpb.GetJoinersResponse{Applicants: applicants}, err
}

//EXPERIMENTAL

//GetJoinersHTTP is an http wrapper for the signup function
func (s *Server) GetJoinersHTTP(w http.ResponseWriter, r *http.Request) {

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

	req := &backendpb.GetJoinersRequest{UUID: uid}

	if strings.Contains(r.Header.Get("Content-type"), "application/json") {
		r.Header.Set("Content-Type", "application/json")
		json.NewDecoder(r.Body).Decode(req)
	} else {
		r.ParseForm()
		req.ProjectID = r.Form.Get("username")
	}

	resp, err := s.GetJoiners(r.Context(), req)
	if err != nil {
		logger.Println("UUID: "+req.UUID+" | err: ", err)
	}

	data := packageResponse(resp, err)

	fmt.Fprintln(w, string(data))

}
