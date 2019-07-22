package sqlstore

import "context"

//Init squishes all the strings together
func (dbConn *SQLStore) Init(ctx context.Context) error {

	_, err := dbConn.db.QueryContext(ctx, `CREATE TABLE IF NOT EXISTS `+dbConn.uploaderTableName+` (

	URL text NOT NULL,
	UploaderID text NOT NULL,
	Base64 text NOT NULL,
	);
`)
}
