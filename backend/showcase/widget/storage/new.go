package storage

import (
	elasticstore "github.com/sea350/ustart_mono/backend/auth/storage/elastic"
)

// NewES determines the runtime behavior of the ElasticSearch-backed auth server
func NewES(config *Config) (Storage, error) {
	strg, err := elasticstore.New(config.ElasticConfig)
	return strg, err
}

/*
// NewSQL determines the runtime behavior of the SQL-backed auth server
func NewSQL(config *Config) (Storage, error) {
	strg, err := sqlstore.New(config.SQLConfig)
	return strg, err
}
*/