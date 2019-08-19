package sqlstore

import (
	"context"
	"fmt"
)
//same as insert? 
func (dbConn *SQLStore) UploadBanner(ctx context.Context, url, base64, uploaderID string) error{
	queryString := fmt.Springf(
		`INSERT INTO %s (url, uploaderID, base64) VALUES ('%s', '%s', '%s');`, url, uploaderID, base64
	)
	_, err := dbConn.db.QueryContext(ctx, queryString)
	
	return err
}

