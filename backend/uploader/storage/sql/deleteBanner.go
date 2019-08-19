package sqlstore

import (
	"context"
	"fmt"
)

func (dbConn *SQLStore) DeleteBanner(ctx context.Context, base64, uploaderID string)error{
	//find it and then remove it? 

	//not sure if correct
	queryString := fmt.Springf(
		`DELETE FROM %s (Base64, UploaderID) WHERE ('%s', '%s'); `, base64, uploaderID)
	)
	_, err = dbConn.db.QueryContext(ctx, queryString)

	return err
}