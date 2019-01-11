package store

// Config is the configuration for the elasticsearch storage client
type Config struct {
	ElasticAddr string
	EIndex      string
	EType       string
}

// NewConfig creates a default config struct
func NewConfig() *Config {
	return &Config{
		ElasticAddr: "localhost:9200",
		EIndex:      "test-auth_data",
		EType:       "AUTH",
	}
}
