package social

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/sea350/ustart_micro/backend/social/socialpb"
)

// LookupBlocked wraps backend/social/lookupblocked.go
func (rapi *RESTAPI) LookupBlocked(w http.ResponseWriter, req *http.Request) {
	regCtx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	req.ParseForm()
	uuid := req.Form.Get("uuid")
	blockedType := req.Form.Get("blockedType")

	lookupblockedReq := &socialpb.LookupBlockedRequest{
		BlockerUUID: uuid,

		BlockedType: blockedType,
	}

	ret := make(map[string]interface{})

	resp, err := rapi.social.LookupBlocked(regCtx, lookupblockedReq)
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
