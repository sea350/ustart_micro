package social

import (
	"github.com/sea350/ustart_micro/backend/social/socialpb"
	"time"
	"context"

)

func (social *Social) Follow(ctx context.Context, req *socialpb.FollowRequest)(*socialpb.FollowResponse, error){

		followerID := req.FollowerUUID
		followerType := req.FollowerType
		followedID :=  req.FollowedUUID
		followedType := req.FollowedType
		notify := req.Notify
		time :=  time.Now()
		sqlStatement := `INSERT INTO follow (follower, followerrtype, followed, followedtype, timestamps) VALUES (" `+ followederID + `","` + followedType+` " ,"` + followedID+` " ,"` + followedType+` " ,"` + time+ `" )`
		_, err = social.sqlClient.Exec(sqlStatement);
		if err != nil {
  		panic(err)
}

		
	
}