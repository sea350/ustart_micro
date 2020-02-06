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
		PID:         req.UUID,
		CustomURL:   req.CustomURL,
		Name:        req.ProjectName,
		Description: "",
		School:      "",
	})
	if err != nil {
		return nil, err
	}
	if resProj == nil {
		logger.Println("this shouldnt happen")
		return nil, errNilResponse
	}

	resUser, err := (*s.profileClient).AddProject(ctx, &profilepb.AddProjectRequest{
		UUID: req.UUID,
		//		ProjectID:resProj.
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

//EXPERIMENTAL

//ProjectNewHTTP is an http wrapper for the new project function
func (s *Server) ProjectNewHTTP(w http.ResponseWriter, r *http.Request) {

	if !setCORS(&w, r) {
		return
	}

	req := &backendpb.SignupRequest{}

	if strings.Contains(r.Header.Get("Content-type"), "application/json") {
		r.Header.Set("Content-Type", "application/json")
		json.NewDecoder(r.Body).Decode(req)
	} else {
		r.ParseForm()
		req.Email = r.Form.Get("email")
		req.Password = r.Form.Get("password")
		req.Username = r.Form.Get("username")
		req.FirstName = r.Form.Get("firstname")
		req.LastName = r.Form.Get("lastname")
	}

	resp, err := s.Signup(r.Context(), req)
	if err != nil {
		logger.Println("Email: "+req.Email+" | err: ", err)
	}

	data := packageResponse(resp, err)

	fmt.Fprintln(w, string(data))

}
