package storage

import (
	"context"
)

// Storage is a database-agnostic interface for persisting showcase/widget data
type Storage interface {
	StoreWidget(context.Context, int, string, string, []string) error
	DeleteWidget(context.Context, string) error
	UpdateBody(context.Context, string, string)
	UpdateIndex(context.Context, string, int)

	//TODO:
	//edit passes in widget ID, field,
	//EditWidget(context.Context, string, string) error
	//LookupWidgetByReferenceID(context.Context, string) (string, error)
	//LookupWidgetByClassificationID(context.Context, string) (string, error)

}