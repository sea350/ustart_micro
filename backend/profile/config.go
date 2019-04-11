package profile

import "github.com/sea350/ustart_micro/backend/profile/storage"

// Config determines the runtime behavior of the Elastic-backed profile server
type Config struct {
	useDummy      bool
	StorageConfig *storage.Config
	DefaultAvatar string
	DefaultBanner string
}
