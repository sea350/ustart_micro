package sqlstore

// Config is the configuration for the elasticsearch storage client
type Config struct {
	DriverName        string
	Port              string
	Host              string
	URL               string
	UploaderID        string
	Base64            string
	UploaderTableName string
}

// NewConfig creates a default config struct
func NewConfig() *Config {
	return &Config{
		DriverName:        "",
		Port:              "",
		Host:              "",
		URL:               "",
		UploaderID:        "",
		Base64:            "",
		UploaderTableName: "",
	}
}
