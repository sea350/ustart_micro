package elasticstore

import (
	"context"
	"time"

	"github.com/olivere/elastic"
)

// ElasticStore implements the storage interface for the Auth module
type ElasticStore struct {
	client *elastic.Client
	eIndex string
	eType  string
}

// New returns a new Eclient elasticstore service
func New(cfg *Config) (*ElasticStore, error) {
	client, err := elastic.NewClient(elastic.SetURL(cfg.ElasticAddr))
	if err != nil {
		return nil, err
	}

	ecl := &ElasticStore{
		client: client,
		eIndex: cfg.EIndex,
		eType:  cfg.EType,
	}

	pingCtx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	_, _, err = ecl.client.Ping(cfg.ElasticAddr).Do(pingCtx)

	return ecl, err
}
