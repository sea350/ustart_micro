package sqlstore

import "context"

// Init initiallizes all required tables
func (dbConn *SQLStore) Init(ctx context.Context) error {
	_, err := dbConn.db.QueryContext(ctx, `CREATE TABLE IF NOT EXISTS  `+dbConn.notifTableName+` (
	notif_id int NOT NULL UNIQUE AUTO_INCREMENT,
	uuid text NOT NULL,
	activity_id text NOT NULL,
	seen bool,
	time_seen timestamp,
	PRIMARY KEY (notif_id)
	);
`)

	return err
}

//CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci
