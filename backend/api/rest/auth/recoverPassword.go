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

// RecoverPassword wraps backend/auth/recoverPassword.go
func (rapi *RESTAPI) RecoverPassword(w http.ResponseWriter, req *http.Request) {
	regCtx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	if !setCORS(&w, req) {
		return
	}

	authReq := &authpb.RecoverRequest{}

	if strings.Contains(req.Header.Get("Content-type"), "application/json") {
		req.Header.Set("Content-Type", "application/json")
		json.NewDecoder(req.Body).Decode(authReq)
	} else {
		req.ParseForm()
		authReq.Email = req.Form.Get("email")
		authReq.Token = req.Form.Get("token")
		authReq.NewPassword = req.Form.Get("newPassword")
	}

	ret := make(map[string]interface{})

	resp, err := rapi.auth.RecoverPassword(regCtx, authReq)
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
		logger.Println("Problem marshaling return data", err)
	}

	fmt.Fprintln(w, string(data))
}
