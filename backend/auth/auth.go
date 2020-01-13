package auth

import (
	"time"

	"github.com/sea350/ustart_micro/backend/auth/storage"
)

const ()

//Auth is an implementation of the auth service defined in service.proto
type Auth struct {
	strg            storage.Storage
	tokenExpiration int
	timeFormat      string
}

// New returns a new SQL auth service
func New(cfg *Config) (*Auth, error) {
	strg, err := storage.NewSQL(cfg.StorageConfig)

	auth := &Auth{
		strg:            strg,
		tokenExpiration: cfg.TokenExpirationHrs,
		timeFormat:      time.RFC3339,
	}

	return auth, err
}
