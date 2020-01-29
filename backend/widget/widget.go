package widget

import (
	"github.com/sea350/ustart_micro/backend/widget/storage"
)

//Widget is an implementation of the widget service as defined in service.proto
type Widget struct {
	strg storage.Storage
}

// New returns a new Eclient widget service
func New(cfg *Config) (*Widget, error) {
	// if cfg.useDummy

	storg, err := storage.NewES(cfg.StorageConfig)

	wid := &Widget{
		strg: storg,
	}

	return wid, err
}
