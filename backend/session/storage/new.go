package storage

import (
	sqlstore "github.com/sea350/ustart_micro/backend/session/storage/sql"
)

// NewSQL determines the runtime behavior of the SQL-backed session manager
func NewSQL(config *Config) (Storage, error) {
	strg, err := sqlstore.New(config.SQLConfig)
	return strg, err
}
