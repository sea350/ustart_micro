package storage

import (
	"context"

	"github.com/sea350/ustart_micro/backend/project/projectpb"
)

// Storage is a database-agnostic interface for persisting data
type Storage interface {
	LookupCustomURL(context.Context, string) (projectpb.Project, error)
	Register(context.Context, string, string, string, string, string, string, string, []string, []string, []string) (string, error)
	Lookup(context.Context, string) (projectpb.Project, error)
	Search(context.Context, []string, bool, bool, bool, map[string][]string, string) ([]string, error)

	ErrProjectDoesNotExist() error
}
