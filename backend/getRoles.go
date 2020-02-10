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

//GetRoles retrieves role profiles associated with a project, also checks permissions to see info
func (s *Server) GetRoles(ctx context.Context, req *backendpb.GetRolesRequest) (*backendpb.GetRolesResponse, error) {

	appRes, err := (*s.projectClient).GetRoles(ctx, &projectpb.GetRolesRequest{
		ProjectID: req.ProjectID,
		UserID:    req.UUID,
	})
	if err != nil {
		return nil, err
	}
	if appRes == nil {
		logger.Println("this shouldnt happen")
		return nil, errNilResponse
	}

	var roles []*backendpb.Role
	for _, role := range appRes.Roles {
		backRole := convertProjRoleToBackRole(role)
		roles = append(roles, backRole)
	}

	return &backendpb.GetRolesResponse{Roles: roles}, err
}



//GetRolesHTTP is an http wrapper for the signup function
func (s *Server) GetRolesHTTP(w http.ResponseWriter, r *http.Request) {

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

	req := &backendpb.GetRolesRequest{UUID: uid}

	if strings.Contains(r.Header.Get("Content-type"), "application/json") {
		r.Header.Set("Content-Type", "application/json")
		json.NewDecoder(r.Body).Decode(req)
	} else {
		r.ParseForm()
		req.ProjectID = r.Form.Get("projectID")
	}

	resp, err := s.GetRoles(r.Context(), req)
	if err != nil {
		logger.Println("UUID: "+req.UUID+" | err: ", err)
	}

	data := packageResponse(resp, err)

	fmt.Fprintln(w, string(data))

}
