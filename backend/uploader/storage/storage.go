package storage

import (
	"bytes"
	"context"
)

// type Storage interface {
// 	Upload(context.Context, string, string) error
// 	Delete(context.Context, string, string) error
// 	//Err____ to handle errors
//Storage holds the necessary functions
type Storage interface {
	Upload(context.Context, *bytes.Reader, string) (string, error) // takes file and file name | returns url of image
	Delete(context.Context, string) error                         //takes url | removes image
}
