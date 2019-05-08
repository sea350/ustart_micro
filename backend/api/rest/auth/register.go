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

// Register wraps backend/auth/register.go
func (rapi *RESTAPI) Register(w http.ResponseWriter, req *http.Request) {
	regCtx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	authReq := &authpb.RegisterRequest{}

	if strings.Contains(req.Header.Get("Content-type"), "application/json") {
		// req.Header.Set("Content-Type", "application/json")
		// w.WriteHeader(http.StatusCreated)
		json.NewDecoder(req.Body).Decode(authReq)
	} else {
		req.ParseForm()
		authReq.Email = req.Form.Get("email")
		authReq.Password = req.Form.Get("password")
	}

	ret := make(map[string]interface{})

	resp, err := rapi.auth.Register(regCtx, authReq)
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
	json.NewEncoder(w).Encode(string(data))
}
