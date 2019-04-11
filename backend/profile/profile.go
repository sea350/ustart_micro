package profile

import (
	"github.com/sea350/ustart_micro/backend/profile/storage"
)

//Profile is an implementation of the profile service as defined in service.proto
type Profile struct {
	strg          *storage.Storage
	defaultAvatar string
	defaultBanner string
}

// New returns a new Eclient profile service
func New(cfg *Config) (*Profile, error) {
	// if cfg.useDummy
	storg, err := storage.NewES(cfg.StorageConfig)

	prof := &Profile{
		strg: &storg,
	}
	return prof, err
}
