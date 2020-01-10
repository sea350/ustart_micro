package social

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/sea350/ustart_micro/backend/social/socialpb"
)

// LookupFollowers wraps backend/social/lookupfollowers.go
func (rapi *RESTAPI) LookupFollowers(w http.ResponseWriter, req *http.Request) {
	regCtx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	req.ParseForm()
	uuid := req.Form.Get("uuid")
	followerType := req.Form.Get("followerType")

	lookupfollowersReq := &socialpb.LookupFollowersRequest{
		UUID:         uuid,
		FollowerType: followerType,
	}

	ret := make(map[string]interface{})

	resp, err := rapi.social.LookupFollowers(regCtx, lookupfollowersReq)
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
