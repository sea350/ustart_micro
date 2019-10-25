package storage

import (
	"github.com/sea350/ustart_micro/backend/auth/storage"
	sqlstore "github.com/sea350/ustart_micro/backend/uploader/storage/sql"
	awsstore "github.com/sea350/ustart_micro/backend/uploader/storage/aws"
)

// Config determines the runtime behavior of the an either SQL or AWS backed customer server
type Config struct {
	StorageConfig *storage.Storage
	//SQLConfig     *sqlstore.Config
	AWSConfig	*awsstore.Config
}

// SQLNewConfig returns a default config object
func SQLNewConfig() *Config {
	return &Config{SQLConfig: sqlstore.NewConfig()}
}

// AWSNewConfig returns a default config object
func AWSNewConfig() *Config {
	return &Config{AWSConfig: awsstore.NewConfig()}
}
