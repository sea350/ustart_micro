package storage

//import "github.com/sea350/ustart_micro/backend/uploader/storage/sql"
type Config struct {
	useDummy     bool
	SQLNewConfig *sqlstore.Config
}

func SQLNewConfig() *Config {
	return &Config{SQLNewConfig: sqlstore.NewConfig()}
}
