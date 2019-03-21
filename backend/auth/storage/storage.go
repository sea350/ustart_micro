package storage

import (
	"context"
	"time"
)

// Storage is a database-agnostic interface for persisting auth data
type Storage interface {
	ChangePassword(context.Context, string, string) error
	GetPassword(context.Context, string) (string, error)
	Lookup(context.Context, string) (string, error)
	Register(context.Context, string, string, string, string, string, time.Time) error
	IDLookup(context.Context, string) (bool, error)
	GetToken(context.Context, string) (string, time.Time, error)
	SetToken(context.Context, string, string, time.Time) error
	Validate(context.Context, string, bool) error

	// rest of the functions
	ErrUserDoesNotExist() error
	ErrTooManyResults() error
	ErrEmailInUse() error
}
