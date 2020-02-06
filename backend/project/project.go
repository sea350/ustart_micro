package project

import (
	"time"

	"github.com/sea350/ustart_micro/backend/project/auth"
	"github.com/sea350/ustart_micro/backend/project/storage"
)

//Project is an implementation of the project service as defined in service.proto
type Project struct {
	strg          storage.Storage
	auth          *auth.Auth
	timeFormat    string
	defaultAvatar string
	defaultBanner string
}

// New returns a new Eclient profile service
func New(cfg *Config) (*Project, error) {
	// if cfg.useDummy

	storg, err := storage.NewES(cfg.StorageConfig)
	if err != nil {
		return nil, err
	}

	aut, err := auth.New(cfg.AuthConfig)
	if err != nil {
		return nil, err
	}

	proj := &Project{
		strg:          storg,
		auth:          aut,
		defaultAvatar: cfg.DefaultAvatar,
		defaultBanner: cfg.DefaultBanner,
		timeFormat:    time.RFC3339,
	}

	return proj, nil
}
