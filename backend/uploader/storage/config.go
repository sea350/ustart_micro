package storage

import (
	awsstore "github.com/sea350/ustart_micro/backend/uploader/storage/aws"
)

// Config determines the runtime behavior of the an either SQL or AWS backed customer server
type Config struct {
	AWSConfig *awsstore.Config
}

// AWSNewConfig returns a default config object
func AWSNewConfig() *Config {
	return &Config{AWSConfig: awsstore.NewConfig()}
}
