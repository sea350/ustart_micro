package sqlstore

// Config is the configuration for the elasticsearch storage client
type Config struct {
	DriverName     string
	Host           string
	Port           string
	DBName         string
	Username       string
	Password       string
	NotifTableName string
}

// NewConfig creates a default config struct
func NewConfig() *Config {
	return &Config{
		DriverName:     "postgresql",
		Host:           "localhost",
		Port:           "5432",
		DBName:         "ustart_sql",
		Username:       "postgres",
		Password:       "postgres",
		NotifTableName: "notifs",
	}
}
