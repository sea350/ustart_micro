package social

import (
	"context"
	"time"

	"github.com/sea350/ustart_micro/backend/social/socialpb"
)

//Follow *insert comment here*
func (social *Social) Follow(ctx context.Context, req *socialpb.FollowRequest) (*socialpb.FollowResponse, error) {

	followerID := req.FollowerUUID
	followerType := req.FollowerType
	followedID := req.FollowedUUID
	followedType := req.FollowedType
	time := time.Now()
	sqlStatement := `INSERT INTO Follow (follower, followertype, followed, followedtype, time) VALUES (" ` + followerID + `","` + followerType + ` " ,"` + followedID + ` " ,"` + followedType + ` " ,"` + time + `" );`
	_, err = social.sqlClient.Exec(sqlStatement)
	if err != nil {
		panic(err)
	}
	return nil, nil

}
