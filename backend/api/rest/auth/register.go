package auth

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/sea350/ustart_mono/backend/auth/authpb"
)

// Register wraps backend/auth/register.go
func (rapi *ElasticAuthRESTAPI) Register(w http.ResponseWriter, req *http.Request) {
	regCtx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	req.ParseForm()
	email := req.Form.Get("email")
	pass := req.Form.Get("password")

	authReq := &authpb.RegisterRequest{
		Email:    email,
		Password: pass,
	}

	resp, err := rapi.eauth.Register(regCtx, authReq)
	if err != nil {
		json.NewEncoder(w).Encode(struct {
			errMsg error
		}{
			errMsg: err,
		})
		return
	}

	json.NewEncoder(w).Encode(resp)
}
