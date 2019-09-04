package sqlstore

import (
	"context"
)

// MemberLookup looks up if a row exists using a certain uuid and pid
func (dbConn *SQLStore) MemberLookup(ctx context.Context, uuid string, pid string) (bool, error) {
	rows, err := dbConn.db.QueryContext(ctx, `SELECT "uuid" FROM `+dbConn.ProjectDataTN+` WHERE uuid= '`+uuid+`AND pid=`+pid+"';")
	if err != nil {
		return false, err
	}

	defer rows.Close()

	if rows.Next() {
		if err := rows.Scan(&uuid); err != nil {
			return false, err
		}

		if rows.Next() {
			return true, errTooManyResults
		}

		return true, nil
	}
	return false, errUserDoesNotExist
}
