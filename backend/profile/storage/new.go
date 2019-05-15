package storage

import (
	elasticstore "github.com/sea350/ustart_micro/backend/profile/storage/elastic"
)

// NewES determines the runtime behavior of the ElasticSearch-backed profile server
func NewES(config *Config) (Storage, error) {
	return elasticstore.New(config.ElasticConfig)
}
