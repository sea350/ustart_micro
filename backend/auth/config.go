package auth

import "github.com/sea350/ustart_mono/backend/auth/storage"

// Config determines the runtime behavior of the redis-backed auth server
type Config struct {
	useDummy      bool
	StorageConfig *storage.Config
}

// NewConfig returns a default config object
func NewConfig() *Config {
	return &Config{StorageConfig: storage.NewConfig()}
}
