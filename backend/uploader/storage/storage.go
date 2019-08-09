package storage

import (
	"context"
)
// type Storage interface {
// 	Upload(context.Context, string, string) error
// 	Delete(context.Context, string, string) error
// 	//Err____ to handle errors
//Storage holds the necessary functions
type Storage interface {
	UploadBanner(conext.Context, string, string)error
	DeleteBanner(context.Context, string, string)error
	UploadProfile(context.Context, string, string)error
	DeleteProfile(context.context, string, string)error

	ErrImproperImport() error
}
