package session

import "github.com/sea350/ustart_micro/backend/session/storage"

// Config determines the runtime behavior of the SQL-backed session manager
type Config struct {
	StorageConfig *storage.Config
	SessionKey    string
}

// NewConfig returns a default config objects
func NewConfig() *Config {
	return &Config{StorageConfig: storage.SQLNewConfig()}
}
