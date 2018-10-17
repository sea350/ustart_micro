package auth

// Config determines the runtime behavior of the redis-backed auth server
type Config struct {
	ElasticAddr string
}

// NewConfig returns a default config object
func NewConfig() *Config {
	return &Config{
		ElasticAddr: "localhost:9200",
	}
}
