package auth

import (
	"time"

	"github.com/sea350/ustart_micro/backend/project/auth/storage"
)

const ()

//Auth is an implementation of the auth service defined in service.proto
type Auth struct {
	Strg       storage.Storage
	timeFormat string
}

// New returns a new SQL auth service
func New(cfg *Config) (*Auth, error) {
	// if cfg.useDummy
	strg, err := storage.NewSQL(cfg.StorageConfig)

	auth := &Auth{
		Strg:       strg,
		timeFormat: time.RFC3339,
	}
	return auth, err
}
