package storage

import (
	"context"

	"github.com/sea350/ustart_micro/backend/profile/profilepb"
)

// Storage is a database-agnostic interface for persisting profile data
type Storage interface {
	Register(context.Context, string, string, string, string, string, string, string, bool, bool) error
	Lookup(context.Context, string) (string, error)
	Search(context.Context, []string, bool, bool, bool, map[string][]string, string) ([]string, error)
	LookupUsername(context.Context, string) (profilepb.Profile, error)

	// rest of the functions
	ErrUserDoesNotExist() error
	ErrTooManyResults() error
}
