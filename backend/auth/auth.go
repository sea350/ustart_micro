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
	estor, err := elasticstore.New(cfg.ElasticConfig)

	auth := &Auth{
		eclient: estor,
	}
	return auth, err
}
