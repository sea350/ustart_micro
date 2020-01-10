package social

import (
	"context"

	"github.com/sea350/ustart_micro/backend/social/socialpb"
)

//Follow *insert comment here*
func (social *Social) Follow(ctx context.Context, req *socialpb.FollowRequest) (*socialpb.FollowResponse, error) {

	followerID := req.FollowerUUID
	followedID := req.FollowedUUID
	// time := time.Now()

	sqlStatement := `DELETE FROM Follow WHERE Follower =` + followerID + ` AND Followed =` + followedID + `;`

	_, err = social.sqlClient.Exec(sqlStatement)
	if err != nil {
		panic(err)
	}
	return nil, nil

}
