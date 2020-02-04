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

//EXPERIMENTAL

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

	ret := make(map[string]interface{})

	resp, err := s.Logout(r.Context(), req)

	var ip string
	if resp != nil && err == nil {
		ip, _, _ = net.SplitHostPort(r.RemoteAddr)
		err = s.sesh.End(ip, &w, r)
	}

	if resp != nil {
		ret["response"] = resp
	} else {
		ret["response"] = ""
	}
	if err != nil {
		ret["error"] = err.Error()
		logger.Println("Logout identifier: "+ip+" | err: ", err)
	} else {
		ret["error"] = ""
	}

	data, err := json.Marshal(ret)
	if err != nil {
		logger.Println("Problem marshaling return data during Logout: ", err)
	}

	fmt.Fprintln(w, string(data))

}
