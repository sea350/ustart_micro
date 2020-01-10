package elasticstore

import (
	"context"

	"github.com/olivere/elastic"
)

// ElasticStore implements the storage interface for the Profile module
type ElasticStore struct {
	client *elastic.Client
	eIndex string
	eType  string
}

// New returns a new Eclient elasticstore service
func New(cfg *Config) (*ElasticStore, error) {
	client, err := elastic.NewClient(elastic.SetURL(cfg.ElasticAddr))
	if err != nil {
		panic(err)
	}

	ecl := &ElasticStore{
		client: client,
		eIndex: cfg.EIndex,
		eType:  cfg.EType,
	}

	// pingCtx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	// defer cancel()

	_, _, err = ecl.client.Ping(cfg.ElasticAddr).Do(context.Background())
	if err != nil {
		panic(err)
	}

	// make sure necessary incdices exist
	//TODO: figure out if what to do with context
	exists, err := ecl.client.IndexExists(ecl.eIndex).Do(context.Background())
	if err != nil {
		panic(err)
	}

	// if not create them
	if !exists {
		createIndex, err := ecl.client.CreateIndex(ecl.eIndex).BodyString(mapping(cfg.EIndex)).Do(context.Background()) //DONT FORGET TO ADD MAPPTING LATER
		if err != nil {
			panic(err)
		}
		// TODO fix this.
		if !createIndex.Acknowledged {
			panic(err)
		}
	}

	return ecl, err
}
