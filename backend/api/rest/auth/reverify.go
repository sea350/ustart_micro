package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/sea350/ustart_mono/backend/auth/authpb"
)

// Reverify wraps backend/auth/reverify.go
func (rapi *RESTAPI) Reverify(w http.ResponseWriter, req *http.Request) {
	regCtx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	req.ParseForm()
	email := req.Form.Get("email")

	verifReq := &authpb.ReverifyRequest{
		Email: email,
	}

	ret := make(map[string]interface{})

	resp, err := rapi.auth.Reverify(regCtx, verifReq)
	ret["response"] = resp
	ret["error"] = err

	data, err := json.Marshal(ret)
	if err != nil {
		logger.Panic(err)
	}

	fmt.Fprintln(w, string(data))
}
