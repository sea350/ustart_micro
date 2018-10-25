package auth

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/sea350/ustart_mono/backend/auth/authpb"
)

// ChangePassword wraps backend/auth/changePassword.go
func (rapi *RESTAPI) ChangePassword(w http.ResponseWriter, req *http.Request) {
	regCtx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	req.ParseForm()
	email := req.Form.Get("email")
	pass := req.Form.Get("password")
	newPass := req.Form.Get("newPassword")

	chPassReq := &authpb.ChangePasswordRequest{
		Email:       email,
		Challenge:   pass,
		NewPassword: newPass,
	}

	resp, err := rapi.auth.ChangePassword(regCtx, chPassReq)
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