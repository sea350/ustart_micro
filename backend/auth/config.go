package auth

import (
	"github.com/sea350/ustart_mono/backend/auth/storage/elastic"
)

// Config determines the runtime behavior of the redis-backed auth server
type Config struct {
	ElasticConfig *elasticstore.Config
}

// NewConfig returns a default config object
func NewConfig() *Config {
	return &Config{ElasticConfig: elasticstore.NewConfig()}
}
