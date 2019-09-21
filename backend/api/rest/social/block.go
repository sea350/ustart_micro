package social

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/sea350/ustart_micro/backend/social/socialpb"
)

// Block wraps backend/social/block.go
func (rapi *RESTAPI) Block(w http.ResponseWriter, req *http.Request) {
	regCtx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	req.ParseForm()
	blockerUUID := req.Form.Get("blockerID")
	blockedUUID := req.Form.Get("blockedID")
	blockerType := req.Form.Get("blockerType")
	blockedType := req.Form.Get("blockedType")

	blockReq := &socialpb.BlockRequest{
		BlockerUUID: blockerUUID,
		BlockedUUId: blockedUUID,
		BlockerType: blockerType,
		BlockedType: blockedType,
	}

	ret := make(map[string]interface{})

	resp, err := rapi.social.Block(regCtx, blockReq)
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
