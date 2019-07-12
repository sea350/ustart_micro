package authapi

import (
	"strconv"

	"github.com/sea350/ustart_micro/backend/auth"
)

//GRPCAPI is the GRPC API implementation for auth
type GRPCAPI struct {
	auth *auth.Auth
	port string
}

// New creates a new auth api, given the config
func New(cfg *Config) (*GRPCAPI, error) {
	auth, err := auth.New(cfg.AuthCfg)
	if err != nil {
		return nil, err
	}

	return &GRPCAPI{
		auth: auth,
		port: strconv.Itoa(cfg.Port),
	}, nil
}
