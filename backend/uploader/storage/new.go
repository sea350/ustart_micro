package storage

import (
	sqlstore "github.com/sea350/ustart_micro/backend/uploader/storage/sql"
	awsstore "github.com/sea350/ustart_micro/backend/uploader/storage/aws"
)

// NewSQL determines the runtime behavior of the ElasticSearch-backed customer server
func NewSQL(config *Config) (Storage, error) {
//	strg, err := &Config{SQLConfig: sqlstore.NewConfig()}
	strg, err := sqlstore.New(config.SQLConfig)
	return strg, err
}

// NewAWS determines the runtime behavior of the AWS-backed customer server
func NewAWS(config *Config) (Storage, error) {
	//	strg, err := &Config{SQLConfig: sqlstore.NewConfig()}
		strg, err := awsstore.New(config.AWSConfig)
		return strg, err
	}