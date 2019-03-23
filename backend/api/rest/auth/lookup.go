package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/sea350/ustart_mono/backend/auth/authpb"
)

// Lookup wraps backend/auth/lookup.go
func (rapi *RESTAPI) Lookup(w http.ResponseWriter, req *http.Request) {
	regCtx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	req.ParseForm()
	email := req.Form.Get("email")

	lookReq := &authpb.LookupRequest{
		Email: email,
	}

	ret := make(map[string]interface{})

	resp, err := rapi.auth.Lookup(regCtx, lookReq)
	ret["response"] = resp
	ret["error"] = err
	data, err := json.Marshal(ret)
	if err != nil {
		logger.Panic(err)
	}

	fmt.Fprintln(w, string(data))
}
