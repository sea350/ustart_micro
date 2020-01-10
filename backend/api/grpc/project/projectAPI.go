package projecteapi

import (
	"strconv"

	"github.com/sea350/ustart_micro/backend/project"
)

//GRPCAPI is the GRPC API implementation for profile
type GRPCAPI struct {
	proj *project.Project
	port string
}

// New creates a new profile api, given the config
func New(cfg *Config) (*GRPCAPI, error) {
	proj, err := project.New(cfg.ProjectCfg)
	if err != nil {
		return nil, err
	}

	return &GRPCAPI{
		proj: proj,
		port: strconv.Itoa(cfg.Port),
	}, nil
}
