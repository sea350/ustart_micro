package project

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/sea350/ustart_micro/backend/project/projectpb"
)

// Register wraps backend/profile/register.go
func (rapi *RESTAPI) Register(w http.ResponseWriter, req *http.Request) {
	regCtx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	req.ParseForm()
	pid := req.Form.Get("pid")
	url := req.Form.Get("customURL")
	pname := req.Form.Get("projectName")
	desc := req.Form.Get("description")
	school := req.Form.Get("school")

	projReq := &projectpb.RegisterRequest{
		PID:         pid,
		CustomURL:   url,
		Name:        pname,
		Description: desc,
		School:      school,
	}

	ret := make(map[string]interface{})

	resp, err := rapi.proj.Register(regCtx, projReq)
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
