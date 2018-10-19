package auth

import (
	elastic "github.com/olivere/elastic"
)

const (
	eIndex = "test-auth_data"
	eType  = "AUTH"
)

// Eclient is an implementation of the auth service defined in service.proto
type Eclient struct {
	client *elastic.Client
}

// New returns a new Eclient auth server
func New(cfg *Config) (*Eclient, error) {
	client, err := elastic.NewClient(elastic.SetURL(cfg.ElasticAddr))
	if err != nil {
		return nil, err
	}

	ecl := &Eclient{
		client: client,
	}

	return ecl, nil
}
