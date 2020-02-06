package storage

import (
	elasticstore "github.com/sea350/ustart_micro/backend/project/storage/elastic"
)

// Config determines the runtime behavior of the an either SQL or ElasticSearch backed auth server
type Config struct {
	useDummy      bool
	ElasticConfig *elasticstore.Config
}

// ESNewConfig returns a default config object
func ESNewConfig() *Config {
	return &Config{ElasticConfig: elasticstore.NewConfig()}
}
