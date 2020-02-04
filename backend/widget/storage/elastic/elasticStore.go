package elasticstore

import (
	"context"

	"github.com/olivere/elastic"
)

// Store implements the storage interface for the widget submodule
type Store struct {
	client *elastic.Client
	eIndex string
	eType  string
}

// New returns a new Eclient elasticstore service
func New(cfg *Config) (*Store, error) {
	client, err := elastic.NewClient(elastic.SetURL(cfg.ElasticAddr))
	if err != nil {
		panic(err)
	}

	ecl := &Store{
		client: client,
		eIndex: cfg.EIndex,
		eType:  cfg.EType,
	}

	_, _, err = ecl.client.Ping(cfg.ElasticAddr).Do(context.Background())
	if err != nil {
		panic(err)
	}

	//making sure necessary incdices exist
	//TODO: figure out if what to do with context
	exists, err := ecl.client.IndexExists(ecl.eIndex).Do(context.Background())
	if err != nil {
		panic(err)
	}

	//if indices dont exist create them
	if !exists {
		//TODO: fix warning
		//WARNING MAPPING INCOMPLETE
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
