package auth

import (
	elastic "github.com/olivere/elastic"
)

const ()

// ElasticAuth is an implementation of the auth service defined in service.proto
type ElasticAuth struct {
	client *elastic.Client
	eIndex string
	eType  string
}

// New returns a new Eclient auth server
func New(cfg *Config) (*ElasticAuth, error) {
	client, err := elastic.NewClient(elastic.SetURL(cfg.ElasticAddr))
	if err != nil {
		return nil, err
	}

	ecl := &ElasticAuth{
		client: client,
		eIndex: cfg.EIndex,
		eType:  cfg.EType,
	}

	return ecl, nil
}
