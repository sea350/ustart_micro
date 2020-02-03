package storage

import (
	elasticstore "github.com/sea350/ustart_micro/backend/widget/storage/elastic"
)

// NewES determines the runtime behavior of the ElasticSearch-backed widget server
func NewES(config *Config) (Storage, error) {
	strg, err := elasticstore.New(config.ElasticConfig)
	return strg, err
}

/*
// NewSQL determines the runtime behavior of the SQL-backed widget server
func NewSQL(config *Config) (Storage, error) {
	strg, err := sqlstore.New(config.SQLConfig)
	return strg, err
}
*/
