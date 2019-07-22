package sqlstore

import (
	"context"
	"fmt"
)

//Insert creates a new row for the SQL table
func (dbConn *SQLStore) Insert(ctx context.Context, url, uploaderID, base64 string) error {
	queryString := fmt.Sprintf(
		"INSERT INTO %s (url, uploaderID, base64) VALUES ('%s', '%s', '%s');",
		url, uploaderID, base64)

	_, err := dbConn.db.QueryContext(ctx, queryString)

	return err
}
