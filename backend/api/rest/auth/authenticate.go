package auth

import (
	"context"
	"encoding/json"
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

	resp, err := rapi.auth.Authenticate(regCtx, authReq)
	if err != nil {
		logger.Println(err)
		err = json.NewEncoder(w).Encode(struct {
			errMsg error
		}{
			errMsg: err,
		})
		if err != nil {
			logger.Println(err)
		}
		return
	}

	json.NewEncoder(w).Encode(resp)
}
