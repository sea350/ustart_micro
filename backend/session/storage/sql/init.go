package sqlstore

import "context"

// Init initiallizes all required tables
func (dbConn *SQLStore) Init(ctx context.Context) error {
	_, err := dbConn.db.QueryContext(ctx, `CREATE TABLE IF NOT EXISTS  `+dbConn.SessionTableName+` (
	session_id int NOT NULL UNIQUE AUTO_INCREMENT,
	uuid text NOT NULL UNIQUE,
	ip_address text,
	token text,
	active_since text,
	expiration text,
	active bool,
	PRIMARY KEY  (sessionID)
	);
`)

	return err
}

//CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci
