package uploader

//Config contains all information needed to run uploader
import "github.com/sea350/ustart_micro/backend/uploader/storage"

//Config struct stores storage layer data
type Config struct {
	StorageConfig *storage.Config
}
