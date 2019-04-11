package storage

import (
	"context"
)

// Storage is a database-agnostic interface for persisting profile data
type Storage interface {
	Register(context.Context, string, string, string, string, string, string, string, bool, bool) error
	Lookup(context.Context, string) (string, error)

	// rest of the functions
	ErrUserDoesNotExist() error
	ErrTooManyResults() error
}
