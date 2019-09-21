package social

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/sea350/ustart_micro/backend/social/socialpb"
)

// CountFollowing wraps backend/social/Countfollowing.go
func (rapi *RESTAPI) CountFollowing(w http.ResponseWriter, req *http.Request) {
	regCtx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	req.ParseForm()
	uuid := req.Form.Get("uuid")
	followingType := req.Form.Get("followingType")

	countfollowingReq := &socialpb.CountFollowingRequest{
		UUID:          uuid,
		FollowingType: followingType,
	}

	ret := make(map[string]interface{})

	resp, err := rapi.social.CountFollowing(regCtx, countfollowingReq)
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
