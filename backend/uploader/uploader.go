package uploader

import (
	"github.com/sea350/ustart_micro/backend/uploader/storage"
)

//Uploader is the central uploader microserves object
type Uploader struct {
	stor storage.Storage
}

//New creates a new Uploader based on the inserted config
func New(cfg *Config) (*Uploader, error) {
	// if cfg.useDummy
	storer, err := storage.NewAWS(cfg.StorageConfig)

	uploader := &Uploader{stor: storer}

	return uploader, err
}
