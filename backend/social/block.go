package social

import (
	"context"
	"time"

	"github.com/sea350/ustart_micro/backend/social/socialpb"
)

//Block ... adds block relationship to table
func (social *Social) Block(ctx context.Context, req *socialpb.BlockRequest) (*socialpb.BlockResponse, error) {

	BlockerID := req.BlockerUUID
	BlockerType := req.BlockerType
	BlockedID := req.BlockedUUID
	BlockedType := req.BlockedType
	time := time.Now()
	sqlStatement := `INSERT INTO Block (Blocker, Blockertype, Blocked, Blockedtype, time) VALUES (" ` + BlockerID + `","` + BlockerType + ` " ,"` + BlockedID + ` " ,"` + BlockedType + ` " ,"` + time + `" );`
	_, err = social.sqlClient.Exec(sqlStatement)
	if err != nil {
		panic(err)
	}
	return &socialpb.BlockResponse{}, nil

}
