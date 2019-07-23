package sqlstore

// Config is the configuration for the elasticsearch storage client
type Config struct {
	URL               string
	UploaderID        string
	Base64            string
	UploaderTableName string
}

// NewConfig creates a default config struct
func NewConfig() *Config {
	return &Config{
		URL:               "",
		UploaderID:        "",
		Base64:            "",
		UploaderTableName: "",
	}
}
