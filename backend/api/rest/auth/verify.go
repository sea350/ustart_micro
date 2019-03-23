package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/sea350/ustart_mono/backend/auth/authpb"
)

// Verify wraps backend/auth/verify.go
func (rapi *RESTAPI) Verify(w http.ResponseWriter, req *http.Request) {
	regCtx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	req.ParseForm()
	email := req.Form.Get("email")
	token := req.Form.Get("token")

	verifReq := &authpb.VerifyRequest{
		Email: email,
		Token: token,
	}

	ret := make(map[string]interface{})

	resp, err := rapi.auth.Verify(regCtx, verifReq)
	ret["response"] = resp
	ret["error"] = err

	data, err := json.Marshal(ret)
	if err != nil {
		logger.Panic(err)
	}

	fmt.Fprintln(w, string(data))
}