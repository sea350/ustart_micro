package project

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/sea350/ustart_micro/backend/project/projectpb"
)

// Lite wraps backend/project/lite.go
func (rapi *RESTAPI) Lite(w http.ResponseWriter, req *http.Request) {
	regCtx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	req.ParseForm()
	pid := req.Form.Get("pid")

	lookReq := &projectpb.LiteRequest{
		ProjectID: pid,
	}

	ret := make(map[string]interface{})

	resp, err := rapi.proj.Lite(regCtx, lookReq)
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
