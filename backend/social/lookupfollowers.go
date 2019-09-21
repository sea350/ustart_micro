package social

import (
	"context"
	"time"

	"github.com/sea350/ustart_micro/backend/social/socialpb"
)

//LookupFollowers ... returns list of UUIDs user is following
func (social *Social) LookupFollowers(ctx context.Context, req *socialpb.FollowRequest) (*socialpb.LookupFollowersResponse, error) {

	uuid := req.UUID
	followerType := req.FollowerType
	time := time.Now()

	sqlStatement := `SELECT followed FROM Follow WHERE follower =` + uuid + `AND followertype =` + followertype + `;`

	res, err = social.sqlClient.Exec(sqlStatement)
	if err != nil {
		panic(err)
	}
	return res, nil

}
