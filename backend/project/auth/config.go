package auth

import "github.com/sea350/ustart_micro/backend/project/auth/storage"

// Config determines the runtime behavior of the SQL-backed auth server
type Config struct {
	StorageConfig *storage.Config
}

// NewConfig returns a default config objects
func NewConfig() *Config {
	return &Config{StorageConfig: storage.SQLNewConfig()}
}
