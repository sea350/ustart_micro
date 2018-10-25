package auth

import "github.com/sea350/ustart_mono/backend/auth"

// GRPCAPI implements a gRPC api
// as a wrapper around the auth package.
type GRPCAPI struct {
	auth *auth.Auth
}

// New creates a new auth api, given the config
func New(cfg *Config) (*GRPCAPI, error) {
	auth, err := auth.New(cfg.AuthCfg)
	if err != nil {
		return nil, err
	}

	return &GRPCAPI{
		auth: auth,
	}, nil
}
