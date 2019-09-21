package social

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/sea350/ustart_micro/backend/social/socialpb"
)

// Unblock wraps backend/social/unblock.go
func (rapi *RESTAPI) Block(w http.ResponseWriter, req *http.Request) {
	regCtx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	req.ParseForm()
	blockerUUID := req.Form.Get("blockerID")
	blockedUUID := req.Form.Get("blockedID")

	unblockReq := &socialpb.UnblockRequest{
		BlockerUUID: blockerUUID,
		BlockedUUID: blockedUUID,
	}

	ret := make(map[string]interface{})

	resp, err := rapi.social.Unblock(regCtx, unblockReq)
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
