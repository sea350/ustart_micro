package auth

import (
	elastic "github.com/olivere/elastic"
)

const (
	eIndex = "test-auth_data"
	eType  = "AUTH"
)

// ElasticAuth is an implementation of the auth service defined in service.proto
type ElasticAuth struct {
	client *elastic.Client
}

// New returns a new Eclient auth server
func New(cfg *Config) (*ElasticAuth, error) {
	client, err := elastic.NewClient(elastic.SetURL(cfg.ElasticAddr))
	if err != nil {
		return nil, err
	}

	ecl := &ElasticAuth{
		client: client,
	}

	return ecl, nil
}
