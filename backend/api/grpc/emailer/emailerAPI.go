package emailerapi

import (
	"strconv"

	"github.com/sea350/ustart_micro/backend/emailer"
)

type GRPCAPI struct {
	emailer *emailer.Emailer
	port    string
}

func New(cfg *Config) (*GRPCAPI, error) {
	emailer, err := emailer.New(cfg.EmailerCfg)
	if err != nil {
		return nil, err
	}

	return &GRPCAPI{
		emailer: emailer,
		port:    strconv.Itoa(cfg.Port),
	}, nil
}
