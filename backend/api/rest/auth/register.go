package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sea350/ustart_mono/backend/auth/authpb"
)

// Register wraps backend/auth/register.go
func (rapi *RESTAPI) Register(w http.ResponseWriter, req *http.Request) {
	// regCtx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	// defer cancel()

	req.ParseForm()
	email := req.Form.Get("email")
	pass := req.Form.Get("password")

	authReq := &authpb.RegisterRequest{
		Email:    email,
		Password: pass,
	}

	ret := make(map[string]interface{})

	resp, err := rapi.auth.Register(context.Background(), authReq)
	if err != nil {
		ret["errMsg"] = err.Error()

		//debugging
		// logger.Panic(err)
	} else {
		ret["response"] = resp
	}

	data, err := json.Marshal(ret)
	if err != nil {
		logger.Panic(err)
	}

	fmt.Fprintln(w, string(data))
}
