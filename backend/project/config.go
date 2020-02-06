package project

import "github.com/sea350/ustart_micro/backend/project/storage"

// Config determines the runtime behavior of the Elastic-backed project server
type Config struct {
	useDummy      bool
	StorageConfig *storage.Config
	DefaultAvatar string
	DefaultBanner string
}
