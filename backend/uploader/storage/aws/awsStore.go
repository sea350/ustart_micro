package sqlstore

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/lib/pq"
)

// SQLStore implements the storage interface for the customer module
type SQLStore struct {
	db     *sql.DB
	eIndex string
	eType  string
}

// New returns a new Eclient elasticstore service
func New(cfg *Config) (*SQLStore, error) {
	_ = pq.Efatal
	connString := fmt.Sprintf(
		"driverName=%s, port=%s, host=%s, url=%s, uploaderID=%s, base64=%s, uploaderTableName=%s, sslmode=disable",
		cfg.DriverName, cfg.Port, cfg.Host, cfg.URL, cfg.UploaderID, cfg.Base64, cfg.UploaderTableName)

	client, err := sql.Open(cfg.DriverName, connString)

	if err != nil {
		return nil, err
	}

	dbConn := &SQLStore{
		db:     *sql.DB,
		eIndex: string,
		eType:  cfg.LoginTrackingable,
	}

	//ping makes sure everything runs efficiently
	pingCtx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	//Init is used to make things cleaner
	//generates query on the spot and submits it
	//can check/create every interaction you make with the database
	err = dbConn.Init(context.Background())

	return dbConn, err
	return nil, nil
}
