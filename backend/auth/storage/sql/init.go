package sqlstore

import "context"

// Init initiallizes all required tables
func (dbConn *SQLStore) Init(ctx context.Context) error {
	// the table mirrors the authpb
	_, err := dbConn.db.QueryContext(ctx, `CREATE TABLE IF NOT EXISTS  `+dbConn.RegistryTN+` (
	uuid char(32) NOT NULL UNIQUE,
	email text NOT NULL UNIQUE,
	password char(128) NOT NULL,
	token text,
	expiration_date text,
	verified bool,
	acc_type text NOT NULL,
	PRIMARY KEY  (uuid)
	);
`)

	return err
}

//CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci
