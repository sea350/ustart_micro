package project

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/sea350/ustart_micro/backend/project/projectpb"
)

// Pull wraps backend/project/pull.go
func (rapi *RESTAPI) Pull(w http.ResponseWriter, req *http.Request) {
	regCtx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	req.ParseForm()
	url := req.Form.Get("customURL")

	lookReq := &projectpb.PullRequest{
		CustomURL: url,
	}

	ret := make(map[string]interface{})

	resp, err := rapi.proj.Pull(regCtx, lookReq)
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
