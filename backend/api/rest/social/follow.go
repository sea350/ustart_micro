package social

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/sea350/ustart_micro/backend/social/socialpb"
)

// Follow wraps backend/social/follow.go
func (rapi *RESTAPI) Follow(w http.ResponseWriter, req *http.Request) {
	regCtx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	req.ParseForm()
	followerUUID := req.Form.Get("followerID")
	followedUUID := req.Form.Get("followedID")
	followerType := req.Form.Get("followerType")
	followedType := req.Form.Get("followedType")

	followReq := &socialpb.FollowRequest{
		FollowerUUID: followerUUID,
		FollowedUUId: followedUUID,
		FollowerType: followerType,
		FollowedType: followedType,
	}

	ret := make(map[string]interface{})

	resp, err := rapi.social.Follow(regCtx, followReq)
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
