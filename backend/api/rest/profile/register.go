package profile

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/sea350/ustart_micro/backend/profile/profilepb"
)

// Register wraps backend/profile/register.go
func (rapi *RESTAPI) Register(w http.ResponseWriter, req *http.Request) {
	regCtx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	req.ParseForm()
	uuid := req.Form.Get("uuid")
	username := req.Form.Get("username")
	fname := req.Form.Get("firstname")
	lname := req.Form.Get("lastname")

	profReq := &profilepb.RegisterRequest{
		UUID:      uuid,
		Username:  username,
		FirstName: fname,
		LastName:  lname,
	}

	ret := make(map[string]interface{})

	resp, err := rapi.prof.Register(regCtx, profReq)
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
