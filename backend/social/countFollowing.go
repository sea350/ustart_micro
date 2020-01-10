package social

import (
	"context"
	"time"

	"github.com/sea350/ustart_micro/backend/social/socialpb"
)

//CountFollowing ... returns number of UUIDs user is following
func (social *Social) CountFollowing(ctx context.Context, req *socialpb.FollowRequest) (*socialpb.CountFollowingResponse, error) {

	uuid := req.UUID
	followerType := req.FollowerType
	time := time.Now()

	sqlStatement := `SELECT COUNT(*) FROM Follow WHERE follower=` + uuid + `AND followertype=` + followerType + `;`
	res, err = social.sqlClient.Exec(sqlStatement)
	if err != nil {
		panic(err)
	}
	return res, nil

}
