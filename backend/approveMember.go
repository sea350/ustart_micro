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

//ApproveMember checks all prerequisite conditions and then adds a new candidate to a project
func (s *Server) ApproveMember(ctx context.Context, req *backendpb.ApproveMemberRequest) (*backendpb.ApproveMemberResponse, error) {

	appRes, err := (*s.projectClient).AddMember(ctx, &projectpb.AddMemberRequest{
		ProjectID:   req.ProjectID,
		AdderID:     req.UUID,
		NewMemberID: req.NewMember,
		RoleName:    req.Role,
	})
	if err != nil {
		return nil, err
	}
	if appRes == nil {
		logger.Println("this shouldnt happen")
		return nil, errNilResponse
	}

	return &backendpb.ApproveMemberResponse{}, err
}

//ApproveMemberHTTP is an http wrapper for the signup function
func (s *Server) ApproveMemberHTTP(w http.ResponseWriter, r *http.Request) {

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

	req := &backendpb.ApproveMemberRequest{UUID: uid}

	if strings.Contains(r.Header.Get("Content-type"), "application/json") {
		r.Header.Set("Content-Type", "application/json")
		json.NewDecoder(r.Body).Decode(req)
	} else {
		r.ParseForm()
		req.ProjectID = r.Form.Get("projectID")
		req.NewMember = r.Form.Get("memberID")
		req.Role = r.Form.Get("role")
	}

	resp, err := s.ApproveMember(r.Context(), req)
	if err != nil {
		logger.Println("UUID: "+req.UUID+" | err: ", err)
	}

	data := packageResponse(resp, err)

	fmt.Fprintln(w, string(data))

}
