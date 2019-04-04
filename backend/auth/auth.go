package auth

import (
	"github.com/sea350/ustart_micro/backend/auth/storage"
)

const ()

// Auth is an implementation of the auth service defined in service.proto
type Auth struct {
	strg            storage.Storage
	tokenExpiration int
}

// New returns a new Eclient auth service
func New(cfg *Config) (*Auth, error) {
	// if cfg.useDummy
	strg, err := storage.NewSQL(cfg.StorageConfig)

	auth := &Auth{
		strg:            strg,
		tokenExpiration: cfg.TokenExpirationHrs,
	}
	return auth, err
}
