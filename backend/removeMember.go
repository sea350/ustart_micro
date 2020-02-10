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

//RemoveMember checks all prerequisite conditions and then removes a member of a project
func (s *Server) RemoveMember(ctx context.Context, req *backendpb.RemoveMemberRequest) (*backendpb.RemoveMemberResponse, error) {

	appRes, err := (*s.projectClient).RemoveMember(ctx, &projectpb.RemoveMemberRequest{
		ProjectID: req.ProjectID,
		RemoverID: req.UUID,
		RemovedID: req.RemovedMemberID,
	})
	if err != nil {
		return nil, err
	}
	if appRes == nil {
		logger.Println("this shouldnt happen")
		return nil, errNilResponse
	}

	return &backendpb.RemoveMemberResponse{}, err
}



//RemoveMemberHTTP is an http wrapper for the signup function
func (s *Server) RemoveMemberHTTP(w http.ResponseWriter, r *http.Request) {

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

	req := &backendpb.RemoveMemberRequest{UUID: uid}

	if strings.Contains(r.Header.Get("Content-type"), "application/json") {
		r.Header.Set("Content-Type", "application/json")
		json.NewDecoder(r.Body).Decode(req)
	} else {
		r.ParseForm()
		req.ProjectID = r.Form.Get("projectID")
		req.RemovedMemberID = r.Form.Get("removedMemberID")
	}

	resp, err := s.RemoveMember(r.Context(), req)
	if err != nil {
		logger.Println("UUID: "+req.UUID+" | err: ", err)
	}

	data := packageResponse(resp, err)

	fmt.Fprintln(w, string(data))

}
