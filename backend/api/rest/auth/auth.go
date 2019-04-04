package auth

import (
	"github.com/sea350/ustart_micro/backend/auth"
)

// RESTAPI implements a REST api
// as a wrapper around the auth package.
type RESTAPI struct {
	auth *auth.Auth
}

// New creates a new auth api, given the config
func New(cfg *Config) (*RESTAPI, error) {
	auth, err := auth.New(cfg.AuthCfg)
	if err != nil {
		return nil, err
	}

	return &RESTAPI{
		auth: auth,
	}, nil
}
