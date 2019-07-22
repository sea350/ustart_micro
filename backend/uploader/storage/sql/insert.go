package sqlstore

import (
	"context"
	"fmt"
	"strconv"
)

//Insert creates a new row for the SQL table
func (dbConn *SQLStore) Insert(ctx context.Context, carID, userID, dateStart string, rate int) error {
	queryString := fmt.Sprintf(
		"INSERT INTO %s (carID, userID, dateStart, rate) VALUES ('%s', '%s', '%s', '"+strconv.Itoa(rate)+"');",
		dbConn.eIndex, carID, userID, dateStart, rate)

	_, err := dbConn.db.QueryContext(ctx, queryString)

	return err
}
