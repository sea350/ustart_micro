package storage

// Storage is a database-agnostic interface for persisting auth data
type Storage interface {
	ChangePassword() error
	GetPassword() error
	Lookup() (bool, error)
	Register() error
	New() (*Storage, error)
	// rest of the functions
}

// // New returns a new Eclient auth service
// func New(cfg *Config) (*Storage, error) {
// 	// if cfg.useDummy
// 	strg, err := elasticstore.New(cfg.ElasticConfig)
// 	return stor, err
// }
