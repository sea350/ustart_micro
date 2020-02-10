package backend

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/sea350/ustart_micro/backend/backendpb"
	"github.com/sea350/ustart_micro/backend/project/projectpb"
)

//RejectJoin checks all prerequisite conditions and then adds a new candidate to a project
func (s *Server) RejectJoin(ctx context.Context, req *backendpb.RejectJoinRequest) (*backendpb.RejectJoinResponse, error) {

	appRes, err := (*s.projectClient).Reject(ctx, &projectpb.RejectRequest{
		ProjectID: req.ProjectID,
		RemoverID: req.UUID,
		RemovedID: req.CandidateID,
	})
	if err != nil {
		return nil, err
	}
	if appRes == nil {
		logger.Println("this shouldnt happen")
		return nil, errNilResponse
	}

	return &backendpb.RejectJoinResponse{}, err
}



//RejectJoinHTTP is an http wrapper for the signup function
func (s *Server) RejectJoinHTTP(w http.ResponseWriter, r *http.Request) {

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

	req := &backendpb.RejectJoinRequest{UUID: uid}

	if strings.Contains(r.Header.Get("Content-type"), "application/json") {
		r.Header.Set("Content-Type", "application/json")
		json.NewDecoder(r.Body).Decode(req)
	} else {
		r.ParseForm()
		req.ProjectID = r.Form.Get("projectID")
		req.CandidateID = r.Form.Get("candidateID")
	}

	resp, err := s.RejectJoin(r.Context(), req)
	if err != nil {
		logger.Println("UUID: "+req.UUID+" | err: ", err)
	}

	data := packageResponse(resp, err)

	fmt.Fprintln(w, string(data))

}
