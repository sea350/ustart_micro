package backend

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"strings"

	"github.com/sea350/ustart_micro/backend/backendpb"
)

//Logout is a placeholder
func (s *Server) Logout(ctx context.Context, req *backendpb.LogoutRequest) (*backendpb.LogoutResponse, error) {

	return &backendpb.LogoutResponse{}, nil
}



//LogoutHTTP is an http wrapper for the logout function
func (s *Server) LogoutHTTP(w http.ResponseWriter, r *http.Request) {

	if !setCORS(&w, r) {
		return
	}

	req := &backendpb.LogoutRequest{}

	if strings.Contains(r.Header.Get("Content-type"), "application/json") {
		r.Header.Set("Content-Type", "application/json")
		json.NewDecoder(r.Body).Decode(req)
	} else {
		r.ParseForm()
		req.SessionID = r.Form.Get("sessionID")
	}

	resp, err := s.Logout(r.Context(), req)

	var ip string
	if resp != nil && err == nil {
		ip, _, _ = net.SplitHostPort(r.RemoteAddr)
		err = s.sesh.End(ip, &w, r)
	}

	if err != nil {
		logger.Println("Logout identifier: "+ip+" | err: ", err)
	}

	data := packageResponse(resp, err)

	fmt.Fprintln(w, string(data))

}
