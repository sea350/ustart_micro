package profileapi

import (
	"strconv"

	"github.com/sea350/ustart_micro/backend/profile"
)

//GPRCAPI is the GRPC API implementation for profile
type GPRCAPI struct {
	prof *profile.Profile
	port string
}

// New creates a new profile api, given the config
func New(cfg *Config) (*GPRCAPI, error) {
	prof, err := profile.New(cfg.ProfileCfg)
	if err != nil {
		return nil, err
	}

	return &GPRCAPI{
		prof: prof,
		port: strconv.Itoa(cfg.Port),
	}, nil
}
