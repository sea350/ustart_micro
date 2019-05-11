package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/sea350/ustart_micro/backend/auth/authpb"
)

// Lookup wraps backend/auth/lookup.go
func (rapi *RESTAPI) Lookup(w http.ResponseWriter, req *http.Request) {
	regCtx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	if !setCORS(&w, req) {
		return
	}

	authReq := &authpb.LookupRequest{}

	if strings.Contains(req.Header.Get("Content-type"), "application/json") {
		req.Header.Set("Content-Type", "application/json")
		json.NewDecoder(req.Body).Decode(authReq)
	} else {
		req.ParseForm()
		authReq.Email = req.Form.Get("email")
	}

	ret := make(map[string]interface{})

	resp, err := rapi.auth.Lookup(regCtx, authReq)
	if resp != nil {
		ret["response"] = resp
	} else {
		ret["response"] = ""
	}
	if err != nil {
		ret["error"] = err.Error()
	} else {
		ret["error"] = ""
	}

	data, err := json.Marshal(ret)
	if err != nil {
		logger.Println("Problem martialing return data", err)
	}

	fmt.Fprintln(w, string(data))
}
