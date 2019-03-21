package sqlstore

import (
	"context"
	"fmt"
	"time"
)

//SetToken sets the token and expiration for the acc of the given email
func (dbConn *SQLStore) SetToken(ctx context.Context, email string, token string, expiration time.Time) error {

	expirationdate := expiration.Format(dbConn.TimeFormat) //format may need to be changed

	queryString := fmt.Sprintf(
		`UPDATE %s SET password= '%s', expiration_date = '%s' WHERE email = "%s";`,
		dbConn.RegistryTN, token, expirationdate, email)

	_, err := dbConn.db.QueryContext(ctx, queryString)

	return err

}
