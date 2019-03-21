package sqlstore

import (
	"context"
	"fmt"
	"time"
)

//Register creates a new row for a new registering user
func (dbConn *SQLStore) Register(ctx context.Context, uuid, email, password, token, accountType string, expiration time.Time) error {
	//todo:
	//BUILD DUPLICATE UUID/EMAIL CONTINGENCIES\

	expirationdate := expiration.Format(dbConn.TimeFormat) //format may need to be changed

	queryString := fmt.Sprintf(
		"INSERT INTO %s (uuid, email, password, token, expiration_date, verified, acc_type, creation_date) VALUES ( '%s', '%s', '%s', '%s', '%s', 'false', '%s', '%s');",
		dbConn.RegistryTN, uuid, email, password, token, expirationdate, accountType, time.Now().Format(dbConn.TimeFormat))

	_, err := dbConn.db.QueryContext(ctx, queryString)

	return err

}
