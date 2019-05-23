package authapi

import (
	"strconv"

	"github.com/sea350/ustart_micro/backend/auth"
)

//GPRCAPI is the GRPC API implementation for auth
type GPRCAPI struct {
	auth *auth.Auth
	port string
}

// New creates a new auth api, given the config
func New(cfg *Config) (*GPRCAPI, error) {
	auth, err := auth.New(cfg.AuthCfg)
	if err != nil {
		return nil, err
	}

	return &GPRCAPI{
		auth: auth,
		port: strconv.Itoa(cfg.Port),
	}, nil
}
