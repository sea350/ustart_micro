package profile

import (
	"github.com/sea350/ustart_micro/backend/profile"
)

// RESTAPI implements a REST api
// as a wrapper around the profile package.
type RESTAPI struct {
	prof *profile.Profile
}

// New creates a new profile api, given the config
func New(cfg *Config) (*RESTAPI, error) {
	prof, err := profile.New(cfg.ProfCfg)
	if err != nil {
		return nil, err
	}

	return &RESTAPI{
		prof: prof,
	}, nil
}
