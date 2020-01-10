package social

import (
	"context"

	"github.com/sea350/ustart_micro/backend/social/socialpb"
)

//Block ... adds block relationship to table
func (social *Social) Block(ctx context.Context, req *socialpb.BlockRequest) (*socialpb.BlockResponse, error) {

	blockerID := req.BlockerUUID

	blockedID := req.BlockedUUID

	// time := time.Now()
	//Time in activity log?

	sqlStatement := `DELETE FROM Block WHERE Blocker =` + blockerID + ` AND Blocked =` + blockedID + `;`
	_, err = social.sqlClient.Exec(sqlStatement)
	if err != nil {
		panic(err)
	}
	return nil, nil

}
