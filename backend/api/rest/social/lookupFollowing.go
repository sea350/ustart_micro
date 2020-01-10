package social

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/sea350/ustart_micro/backend/social/socialpb"
)

// LookupFollowing wraps backend/social/lookupfollowing.go
func (rapi *RESTAPI) LookupFollowing(w http.ResponseWriter, req *http.Request) {
	regCtx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	req.ParseForm()
	uuid := req.Form.Get("uuid")
	followingType := req.Form.Get("followingType")

	lookupfollowingReq := &socialpb.LookupFollowingRequest{
		UUID:          uuid,
		FollowingType: followingType,
	}

	ret := make(map[string]interface{})

	resp, err := rapi.social.LookupFollowing(regCtx, lookupfollowingReq)
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
