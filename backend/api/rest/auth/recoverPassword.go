package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/sea350/ustart_mono/backend/auth/authpb"
)

// RecoverPassword wraps backend/auth/recoverPassword.go
func (rapi *RESTAPI) RecoverPassword(w http.ResponseWriter, req *http.Request) {
	regCtx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	req.ParseForm()
	email := req.Form.Get("email")
	token := req.Form.Get("token")
	newPassword := req.Form.Get("newPassword")

	recReq := &authpb.RecoverRequest{
		Email:       email,
		Token:       token,
		NewPassword: newPassword,
	}

	ret := make(map[string]interface{})

	resp, err := rapi.auth.RecoverPassword(regCtx, recReq)
	ret["response"] = resp
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
