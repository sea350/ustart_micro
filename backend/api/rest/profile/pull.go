package profile

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/sea350/ustart_micro/backend/profile/profilepb"
)

// Pull wraps backend/profile/pull.go
func (rapi *RESTAPI) Pull(w http.ResponseWriter, req *http.Request) {
	regCtx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	req.ParseForm()
	username := req.Form.Get("username")

	lookReq := &profilepb.PullRequest{
		Username: username,
	}

	ret := make(map[string]interface{})

	resp, err := rapi.prof.Pull(regCtx, lookReq)
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
