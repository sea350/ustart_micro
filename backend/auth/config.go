package auth

import "github.com/sea350/ustart_micro/backend/auth/storage"

// Config determines the runtime behavior of the SQL-backed auth server
type Config struct {
	useDummy           bool
	TokenExpirationHrs int
	StorageConfig      *storage.Config
}

// NewConfig returns a default config objects
func NewConfig() *Config {
	return &Config{StorageConfig: storage.SQLNewConfig()}
}
