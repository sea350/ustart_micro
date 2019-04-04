package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/sea350/ustart_micro/backend/auth/authpb"
)

// Reverify wraps backend/auth/reverify.go
func (rapi *RESTAPI) Reverify(w http.ResponseWriter, req *http.Request) {
	regCtx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	req.ParseForm()
	email := req.Form.Get("email")

	verifReq := &authpb.ReverifyRequest{
		Email: email,
	}

	ret := make(map[string]interface{})

	resp, err := rapi.auth.Reverify(regCtx, verifReq)
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
