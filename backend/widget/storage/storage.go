package storage

import (
	"context"

	"github.com/sea350/ustart_micro/backend/widget/widgetpb"
)

// Storage is a database-agnostic interface for persisting showcase/widget data
type Storage interface {
	StoreWidget(context.Context, int, string, string, []*widgetpb.Reference) (string, error)
	DeleteWidget(context.Context, string) error
	UpdateBody(context.Context, string, string) error
	UpdateIndex(context.Context, string, int) error
	GetShowcase(context.Context, string) (int, []widgetpb.Widget, error)

	//TODO:

	//LookupWidgetByReferenceID(context.Context, string) (string, error)
	//LookupWidgetByClassificationID(context.Context, string) (string, error)

}
