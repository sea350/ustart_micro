package uploader

import (
	"github.com/sea350/ustart_micro/backend/uploader/storage"
)


//New creates a new Uploader based on the inserted config
func New(cfg *Config) (*Uploader, error) {
	// if cfg.useDummy
	uploader, err := storage.NewAWS(cfg.StorageConfig)
	
	return uploader, err
}
