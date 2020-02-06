package storage

import (
	elasticstore "github.com/sea350/ustart_micro/backend/project/storage/elastic"
)

// NewES determines the runtime behavior of the ElasticSearch-backed auth server
func NewES(config *Config) (Storage, error) {
	strg, err := elasticstore.New(config.ElasticConfig)
	return strg, err
}
