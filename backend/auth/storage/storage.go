package storage

import "context"

// Storage is a database-agnostic interface for persisting auth data
type Storage interface {
	ChangePassword(context.Context, string, string) error
	GetPassword(context.Context, string) (string, error)
	Lookup(context.Context, string) (bool, error)
	Register(context.Context, string, string, string, string) error
	//New(interface{}) (*Storage, error)
	// rest of the functions
}

// // New returns a new Eclient auth service
// func New(cfg *Config) (*Storage, error) {
// 	// if cfg.useDummy
// 	strg, err := elasticstore.New(cfg.ElasticConfig)
// 	return stor, err
// }
