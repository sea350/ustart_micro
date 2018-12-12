package social

import (
	"github.com/sea350/ustart_mono/backend/social/storage"
)

// Config determines the runtime behavior of the redis-backed auth server
type Config struct {
	SqlStoreConfig *sqlstore.Config
}

// NewConfig returns a default config object
func NewConfig() *Config {
	return &Config {SqlStoreConfig: sqlstore.NewConfig()}
}


