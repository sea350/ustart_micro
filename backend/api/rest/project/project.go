package project

import "github.com/sea350/ustart_micro/backend/project"

// RESTAPI implements a REST api
// as a wrapper around the project package.
type RESTAPI struct {
	proj *project.Project
}

// New creates a new project api, given the config
func New(cfg *Config) (*RESTAPI, error) {
	proj, err := project.New(cfg.ProjCfg)
	if err != nil {
		return nil, err
	}

	return &RESTAPI{
		proj: proj,
	}, nil
}
