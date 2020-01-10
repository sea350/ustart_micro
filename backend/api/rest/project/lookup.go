package project

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Lookup wraps backend/profile/lookup.go
func (rapi *RESTAPI) Lookup(w http.ResponseWriter, req *http.Request) {
	regCtx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	req.ParseForm()
	url := req.Form.Get("customURL")

	lookReq := &profilepb.LookupRequest{
		Username: username,
	}

	ret := make(map[string]interface{})

	resp, err := rapi.prof.Lookup(regCtx, lookReq)
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
