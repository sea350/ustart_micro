package storage

import (
	"context"
)

type Storage interface {
	Upload(context.Context, string, string) error
	Delete(context.Context, string, string) error
	//Err____ to handle errors
}
