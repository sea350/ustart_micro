package sqlstore

import (
	"context"
	"database/sql"
	"time"
)

// SQLStore implements the storage interface for the Auth module
type SQLStore struct {
	db              *sql.DB
	RegistryTN      string
	LoginTrackingTN string
}

// New returns a new SQLStore service
func New(cfg *Config) (*SQLStore, error) {
	client, err := sql.Open(cfg.DriverName, "host="+cfg.Host+" port="+cfg.Port+" dbname="+cfg.DBName+" user="+cfg.Username+" password="+cfg.Password+";")
	if err != nil {
		return nil, err
	}

	dbConn := &SQLStore{
		db:              client,
		RegistryTN:      cfg.RegistryTable,
		LoginTrackingTN: cfg.LoginTrackingTable,
	}

	pingCtx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	err = dbConn.db.PingContext(pingCtx)

	return dbConn, err
}
