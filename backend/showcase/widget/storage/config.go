package storage

import (
	elasticstore "github.com/sea350/ustart_micro/backend/showcase/widget/storage/elastic"
	//sqlstore "github.com/sea350/ustart_micro/backend/showcase/widget/storage/sql"
)

// Config determines the runtime behavior of the an either SQL or ElasticSearch backed showcase/widget server
type Config struct {
	useDummy      bool
	ElasticConfig *elasticstore.Config
	//SQLConfig     *sqlstore.Config
}

// ESNewConfig returns a default config object
func ESNewConfig() *Config {
	return &Config{ElasticConfig: elasticstore.NewConfig()}
}

/*
// SQLNewConfig returns a default config object
func SQLNewConfig() *Config {
	return &Config{SQLConfig: sqlstore.NewConfig()}
}
*/
