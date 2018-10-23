package storage

// Storage is a database-agnostic interface for persisting auth data
type Storage interface {
	Register() error
	// rest of the functions
}
