package sqlstore

// Config is the configuration for the elasticsearch storage client
type Config struct {
	Host   string
	Port   string 
	Database string
	Password string
	User   string
}

// NewConfig creates a default config struct
func NewConfig() *Config {
	return &Config{
		Host: "localhost",
		Port: "5432", 
		Database: "ustart",
		Password: "test",
		User: "test",
	
	}
}
