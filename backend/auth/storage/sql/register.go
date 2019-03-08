package sqlstore

import (
	"context"
	"fmt"
)

//Register creates a new row for a new registering user
func (dbConn *SQLStore) Register(ctx context.Context, email string, password, token string, accountType string) error {
	//todo:
	//FIGURE OUT HOW AND WHERE TO GEN UUID
	//IMPORT EXPIRATION DATE
	//BUILD DUPLICATE UUID/EMAIL CONTINGENCIES
	uuid := "test123"
	var expirationdate string
	queryString := fmt.Sprintf(
		"INSERT INTO %s (uuid, email, password, token, expiration_date, verified, acc_type) VALUES ( '%s', '%s', '%s', '%s', '%s', 'false', '%s');",
		dbConn.RegistryTN, uuid, email, password, token, expirationdate, accountType)
	_, err := dbConn.db.QueryContext(ctx, queryString)
	// if err == nil {
	// 	rows.Close()
	// }

	return err

}
