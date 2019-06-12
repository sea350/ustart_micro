package storage

import (
	"context"
)

// Storage is a database-agnostic interface for persisting showcase/widget data
type Storage interface {
	StoreWidget(context.Context, string) error
	DeleteWidget(context.Context, string) error

	//TODO:
	//edit passes in widget ID, field,
	//EditWidget(context.Context, string, string) error
	//LookupWidgetByReferenceID(context.Context, string) (string, error)
	//LookupWidgetByClassificationID(context.Context, string) (string, error)

}
