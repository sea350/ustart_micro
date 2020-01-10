package project

import (
	"github.com/sea350/ustart_micro/backend/project/storage"
)

//Project is an implementation of the project service as defined in service.proto
type Project struct {
	strg          storage.Storage
	defaultAvatar string
	defaultBanner string
}

// New returns a new Eclient profile service
func New(cfg *Config) (*Project, error) {
	// if cfg.useDummy

	storg, err := storage.NewES(cfg.StorageConfig)

	proj := &Project{
		strg:          storg,
		defaultAvatar: cfg.DefaultAvatar,
		defaultBanner: cfg.DefaultBanner,
	}

	return proj, err
}
