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

//ProjectNew registers a new project and updates associated services
func (s *Server) ProjectNew(ctx context.Context, req *backendpb.ProjectNewRequest) (*backendpb.ProjectNewResponse, error) {

	resProj, err := (*s.projectClient).Register(ctx, &projectpb.RegisterRequest{
		CreatorID:   req.UUID,
		CustomURL:   req.CustomURL,
		Name:        req.ProjectName,
		Description: req.Description,
		School:      "",
		Avatar:      req.Avatar,
	})
	if err != nil {
		return nil, err
	}
	if resProj == nil {
		logger.Println("this shouldnt happen")
		return nil, errNilResponse
	}

	resUser, err := (*s.profileClient).AddProject(ctx, &profilepb.AddProjectRequest{
		UUID:      req.UUID,
		ProjectID: resProj.ProjectID,
	})
	if err != nil {
		return nil, err
	}
	if resUser == nil {
		logger.Println("this shouldnt happen")
		return nil, errNilResponse
	}

	return &backendpb.ProjectNewResponse{}, nil
}



//ProjectNewHTTP is an http wrapper for the new project function
func (s *Server) ProjectNewHTTP(w http.ResponseWriter, r *http.Request) {

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

	req := &backendpb.ProjectNewRequest{UUID: uid}

	if strings.Contains(r.Header.Get("Content-type"), "application/json") {
		r.Header.Set("Content-Type", "application/json")
		json.NewDecoder(r.Body).Decode(req)
	} else {
		r.ParseForm()
		req.ProjectName = r.Form.Get("projectName")
		req.CustomURL = r.Form.Get("customURL")
		req.Category = r.Form.Get("category")
		req.Description = r.Form.Get("description")
		req.Avatar = r.Form.Get("avatar")
	}

	resp, err := s.ProjectNew(r.Context(), req)
	if err != nil {
		logger.Println("UUID: "+uid+" | err: ", err)
	}

	data := packageResponse(resp, err)

	fmt.Fprintln(w, string(data))

}
