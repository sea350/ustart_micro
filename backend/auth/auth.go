package auth

import (
	"github.com/sea350/ustart_mono/backend/auth/storage"
)

const ()

// Auth is an implementation of the auth service defined in service.proto
type Auth struct {
	strg storage.Storage
}

// New returns a new Eclient auth service
func New(cfg *Config) (*Auth, error) {
	// if cfg.useDummy
	strg, err := storage.NewSQL(cfg.StorageConfig)

	auth := &Auth{
		strg: strg,
	}
	return auth, err
}
