package storage

import (
	sqlstore "github.com/sea350/ustart_micro/backend/session/storage/sql"
)

// Config determines the runtime behavior of the SQL or backed session manager
type Config struct {
	SQLConfig *sqlstore.Config
}

// SQLNewConfig returns a default config object
func SQLNewConfig() *Config {
	return &Config{SQLConfig: sqlstore.NewConfig()}
}
