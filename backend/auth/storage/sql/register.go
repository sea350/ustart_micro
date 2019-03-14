package sqlstore

import (
	"context"
	"fmt"
	"time"
)

//Register creates a new row for a new registering user
func (dbConn *SQLStore) Register(ctx context.Context, uuid string, email string, password, token string, accountType string, expiration time.Time) error {
	//todo:
	//BUILD DUPLICATE UUID/EMAIL CONTINGENCIES\

	expirationdate := expiration.String() //format may need to be changed

	queryString := fmt.Sprintf(
		"INSERT INTO %s (uuid, email, password, token, expiration_date, verified, acc_type) VALUES ( '%s', '%s', '%s', '%s', '%s', 'false', '%s');",
		dbConn.RegistryTN, uuid, email, password, token, expirationdate, accountType)

	_, err := dbConn.db.QueryContext(ctx, queryString)

	return err

}
