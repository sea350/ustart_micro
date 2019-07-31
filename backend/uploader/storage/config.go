package storage

<<<<<<< HEAD
//import "github.com/sea350/ustart_micro/backend/uploader/storage/sql"
type Config struct {
	useDummy     bool
	SQLNewConfig *sqlstore.Config
}

func SQLNewConfig() *Config {
	return &Config{SQLNewConfig: sqlstore.NewConfig()}
=======
import (
	"github.com/sea350/ustart_micro/backend/auth/storage"
	sqlstore "github.com/sea350/ustart_micro/backend/uploader/storage/sql"
)

// Config determines the runtime behavior of the an either SQL or ElasticSearch backed customer server
type Config struct {
	StorageConfig *storage.Storage
	SQLConfig     *sqlstore.Config
}

// SQLNewConfig returns a default config object
func SQLNewConfig() *Config {
	return &Config{SQLConfig: sqlstore.NewConfig()}
>>>>>>> 6614f1cab1506aae051d86b454c0364da155b22b
}
