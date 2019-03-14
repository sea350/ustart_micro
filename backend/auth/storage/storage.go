package storage

import (
	"context"
	"time"
)

// Storage is a database-agnostic interface for persisting auth data
type Storage interface {
	ChangePassword(context.Context, string, string) error
	GetPassword(context.Context, string) (string, error)
	Lookup(context.Context, string) (bool, error)
	Register(context.Context, string, string, string, string, string, time.Time) error
	IDLookup(context.Context, string) (bool, error)
	//New(interface{}) (*Storage, error)
	// rest of the functions
}
