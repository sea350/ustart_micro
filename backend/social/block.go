package social

import (
	"context"
	"time"
	"github.com/sea350/ustart_mono/backend/social/socialpb"
	//"fmt"

)

func (social *Social) Block(ctx context.Context, req *socialpb.BlockRequest)(*socialpb.BlockResponse, error){

		blockerID := req.BlockerUUID
		blockerType := req.BlockerType
		blockedID :=  req.BlockedUUID
		blockedType := req.BlockedType
		time :=  time.Now()
		sqlStatement := `INSERT INTO block (blocker, blockertype, blocked, blockedtype, timestamps) VALUES (" `+ blockerID + `","` + blockerType+` " ,"` + blockedID+` " ,"` + blockedType+` " ,"` + time+ `" )`
		_, err := social.sqlClient.Exec(sqlStatement);
		if err != nil {
  		panic(err)
		}

		return &socialpb.BlockResponse{}, err

}