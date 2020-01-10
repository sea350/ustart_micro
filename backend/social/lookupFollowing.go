package social

import (
	"context"
	"time"

	"github.com/sea350/ustart_micro/backend/social/socialpb"
)

//LookupFollowing ... returns list of UUIDs following the user
func (social *Social) LookupFollowing(ctx context.Context, req *socialpb.FollowRequest) (*socialpb.LookupFollowingResponse, error) {

	uuid := req.UUID
	followerType := req.FollowerType
	time := time.Now()

	sqlStatement := `SELECT follower FROM Follow WHERE followed =` + uuid + `AND followertype =` + followertype + `;`

	res, err = social.sqlClient.Exec(sqlStatement)
	if err != nil {
		panic(err)
	}
	return res, nil

}
