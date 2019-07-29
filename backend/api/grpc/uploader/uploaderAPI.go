package uploaderapi

import (
	"strconv"

	"github.com/sea350/ustart_micro/backend/uploader"
)

//GRPCAPI is the GRPC API implementation for uploader
type GRPCAPI struct {
	upload *uploader.Uploader
	port   string
}

// New creates a new uploader api, given the config
func New(cfg *Config) (*GRPCAPI, error) {
	upload, err := uploader.New(cfg.UploaderCfg)
	if err != nil {
		return nil, err
	}

	return &GRPCAPI{
		upload: upload,
		port:   strconv.Itoa(cfg.Port),
	}, nil
}
