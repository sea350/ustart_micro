package social

import (
	sqlstore "github.com/sea350/ustart_micro/backend/social/storage"
)

// Config determines the runtime behavior of the redis-backed auth server
type Config struct {
	SQLStoreConfig *sqlstore.Config
}

// NewConfig returns a default config object
func NewConfig() *Config {

	return &Config{SQLStoreConfig: sqlstore.NewConfig()}
}
