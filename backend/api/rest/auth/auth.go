package auth

import (
	"github.com/sea350/ustart_mono/backend/auth"
)

// ElasticAuthRESTAPI implements a REST api
// as a wrapper around the auth package.
type ElasticAuthRESTAPI struct {
	eauth *auth.ElasticAuth
}

// New creates a new auth api, given the config
func New(cfg *Config) (*ElasticAuthRESTAPI, error) {
	eauth, err := auth.New(cfg.EAuthCfg)
	if err != nil {
		return nil, err
	}

	return &ElasticAuthRESTAPI{
		eauth: eauth,
	}, nil
}
