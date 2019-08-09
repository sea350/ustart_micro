package storage

import (
	sqlstore "github.com/sea350/ustart_micro/backend/uploader/storage/sql"
)

// NewSQL determines the runtime behavior of the ElasticSearch-backed customer server
func NewSQL(config *Config) (Storage, error) {
//	strg, err := &Config{SQLConfig: sqlstore.NewConfig()}
	strg, err := sqlstore.New(config.SQLConfig)
	return strg, err
}
