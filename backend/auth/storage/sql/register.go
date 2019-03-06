package sqlstore

import (
	"context"
)

//Register creates a new row for a new registering user
func (dbConn *SQLStore) Register(ctx context.Context, email string, password, token string, accountType string) error {
	//todo:
	//FIGURE OUT HOW AND WHERE TO GEN UUID
	//IMPORT EXPIRATION DATE
	var uuid string
	var expirationdate string
	_, err := dbConn.db.QueryContext(ctx, "INSERT INTO "+dbConn.RegistryTN+" (uuid, email, password, token, expiration_date, verified, acc_type) VALUES ( "+uuid+", "+email+", "+password+", "+token+", "+expirationdate+", false, "+accountType+");")

	return err

}
