package sqlstore

import "context"

// Init initiallizes all required tables
func (dbConn *SQLStore) Init(ctx context.Context) error {
	_, err := dbConn.db.QueryContext(ctx, `CREATE TABLE IF NOT EXISTS  `+dbConn.activityTableName+` (
	activity_id int NOT NULL UNIQUE AUTO_INCREMENT,
	actor_id text NOT NULL UNIQUE,
	object_id text,
	action text,
	time timestamp,
	PRIMARY KEY  (activity_id)
	);
`)

	return err
}

//CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci
