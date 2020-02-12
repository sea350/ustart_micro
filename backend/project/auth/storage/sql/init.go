package sqlstore

import "context"

// Init initiallizes all required tables
func (dbConn *SQLStore) Init(ctx context.Context) error {
	// the table mirrors the authpb
	_, err := dbConn.db.ExecContext(ctx, `CREATE TABLE IF NOT EXISTS  `+dbConn.memberTN+` (
	uuid text NOT NULL,
	project_id text NOT NULL,
	role_name text NOT NULL,
	join_date text,
	PRIMARY KEY (uuid, project_id)
);
`)

	if err != nil {
		return err
	}
	_, err = dbConn.db.ExecContext(ctx, `CREATE TABLE IF NOT EXISTS  `+dbConn.roleTN+` (
	role_name text NOT NULL,
	project_id text NOT NULL,
	manage_members boolean,
	change_icon boolean,
	change_banner boolean,
	manage_entries boolean,
	manage_links boolean,
	manage_tags boolean,
	PRIMARY KEY (role_name, project_id)
);
`)

	// the table mirrors the authpb
	_, err = dbConn.db.ExecContext(ctx, `CREATE TABLE IF NOT EXISTS  `+dbConn.requestTN+` (
	uuid text NOT NULL,
	project_id text NOT NULL,
	request_date text,
	PRIMARY KEY (uuid, project_id)
);
`)
	return err
}

//CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci
