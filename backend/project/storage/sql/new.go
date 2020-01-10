package sqlstore

import (
	sqlstore "github.com/sea350/ustart_micro/backend/project/storage/sql"
)

// NewSQL determines the runtime behavior of the SQL-backed auth server
func NewSQL(config *Config) (Storage, error) {
	strg, err := sqlstore.New(config.SQLConfig)
	return strg, err
}
