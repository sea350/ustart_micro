package sqlstore

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/lib/pq"
)

// SQLStore implements the storage interface for the Auth module
type SQLStore struct {
	db         *sql.DB
	memberTN   string //Table name for project members
	roleTN     string //Table name for project Roles
	requestTN  string //Table name for Join Requests
	TimeFormat string
}

// New returns a new SQLStore service
func New(cfg *Config) (*SQLStore, error) {
	_ = pq.Efatal
	connString := fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		cfg.Username, cfg.Password, cfg.DBName, cfg.Host, cfg.Port)

	client, err := sql.Open(cfg.DriverName, connString)
	if err != nil {
		return nil, err
	}

	dbConn := &SQLStore{
		db:         client,
		memberTN:   cfg.MemberTableName,
		roleTN:     cfg.RoleTableName,
		requestTN:  cfg.RequestTableName,
		TimeFormat: time.RFC3339,
	}

	pingCtx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	err = dbConn.db.PingContext(pingCtx)
	if err != nil {
		return nil, err
	}

	err = dbConn.Init(context.Background())

	return dbConn, err
}
