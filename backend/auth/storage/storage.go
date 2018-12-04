package storage

import (
	"github.com/sea350/ustart_mono/backend/auth/storage/elastic"
)

// Storage is a database-agnostic interface for persisting auth data
type Storage interface {
	Register() error
	GetPassword() error

	// rest of the functions
}

// New returns a new Eclient auth service
func New(cfg *Config) (*Storage, error) {
	// if cfg.useDummy
	strg, err := elasticstore.New(cfg.ElasticConfig)

	stor := &Storage{
		eclient: estor,
	}
	return stor, err
}
