package sqlstore

import "context"

// Init initiallizes all required tables
func (dbConn *SQLStore) Init(ctx context.Context) error {
	// the table mirrors the authpb
	_, err := dbConn.db.QueryContext(ctx, `CREATE TABLE  `+dbConn.RegistryTN+` (
	uuid text NOT NULL UNIQUE,
	email text NOT NULL UNIQUE,
	password text NOT NULL,
	token text,
	expiration_date text,
	verified bool,
	acc_type text NOT NULL,
	PRIMARY KEY  (uuid)
	);
ALTER TABLE
	`+dbConn.RegistryTN+`
	CONVERT TO CHARACTER SET utf8mb4
	COLLATE utf8mb4_unicode_ci;
`)

	return err
}
