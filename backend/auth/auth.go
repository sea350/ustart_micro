package auth

import (
	"github.com/sea350/ustart_mono/backend/auth/storage/elastic"
)

const ()

// Auth is an implementation of the auth service defined in service.proto
type Auth struct {
	eclient *elasticstore.ElasticStore
}

// New returns a new Eclient auth service
func New(cfg *Config) (*Auth, error) {
	conf := &elasticstore.Config{
		ElasticAddr: cfg.ElasticAddr,
		EIndex:      cfg.EIndex,
		EType:       cfg.EType,
	}
	estor, err := elasticstore.New(conf)

	auth := &Auth{
		eclient: estor,
	}
	return auth, err
}
