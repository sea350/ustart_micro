package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/sea350/ustart_mono/backend/auth/authpb"
)

// Authenticate wraps backend/auth/authenticate.go
func (rapi *RESTAPI) Authenticate(w http.ResponseWriter, req *http.Request) {
	regCtx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	req.ParseForm()
	email := req.Form.Get("email")
	pass := req.Form.Get("password")

	authReq := &authpb.AuthenticateRequest{
		Email:     email,
		Challenge: pass,
	}

	ret := make(map[string]interface{})

	resp, err := rapi.auth.Authenticate(regCtx, authReq)
	ret["response"] = resp
	ret["error"] = err
	data, err := json.Marshal(ret)
	if err != nil {
		logger.Panic(err)
	}

	fmt.Fprintln(w, string(data))
}
