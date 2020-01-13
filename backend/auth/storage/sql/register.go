package sqlstore

import (
	"context"
	"fmt"
)

//Register creates a new row for a new registering user
func (dbConn *SQLStore) Register(ctx context.Context, uuid, email, password, token, accountType, creationTime, expiration string) error {
	//todo:
	//BUILD DUPLICATE UUID/EMAIL CONTINGENCIES\

	queryString := fmt.Sprintf(
		"INSERT INTO %s (uuid, email, password, token, expiration_date, verified, acc_type, creation_date) VALUES ( $1, $2, $3, $4, $5, $8, $6, $7);",
		dbConn.registryTN)

	_, err := dbConn.db.QueryContext(ctx, queryString, uuid, email, password, token, expiration, accountType, creationTime, false)

	return err

}
