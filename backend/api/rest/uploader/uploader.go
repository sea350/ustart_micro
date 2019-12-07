package uploader

import (
	"github.com/sea350/ustart_micro/backend/uploader"
)

// RESTAPI implements a REST api
// as a wrapper around the uploader package.
type RESTAPI struct {
	uploader *uploader.Uploader
}

// New creates a new auth api, given the config
func New(cfg *Config) (*RESTAPI, error) {
	uploader, err := uploader.New(cfg.UploaderCfg)
	if err != nil {
		return nil, err
	}

	return &RESTAPI{
		uploader: uploader,
	}, nil
}
