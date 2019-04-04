package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/sea350/ustart_micro/backend/auth/authpb"
)

// ChangePassword wraps backend/auth/changePassword.go
func (rapi *RESTAPI) ChangePassword(w http.ResponseWriter, req *http.Request) {
	regCtx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
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

	ret := make(map[string]interface{})

	resp, err := rapi.auth.ChangePassword(regCtx, chPassReq)
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
		logger.Panic(err)
	}

	fmt.Fprintln(w, string(data))
}
