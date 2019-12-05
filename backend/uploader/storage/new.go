package storage

import (
	awsstore "github.com/sea350/ustart_micro/backend/uploader/storage/aws"
)

// NewAWS determines the runtime behavior of the AWS-backed customer server
func NewAWS(config *Config) (Storage, error) {
	//	strg, err := &Config{SQLConfig: sqlstore.NewConfig()}
	strg, err := awsstore.New(config.AWSConfig)
	return strg, err
}
