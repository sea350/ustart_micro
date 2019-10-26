package uploader

//Config contains all information needed to run uploader
import "github.com/sea350/ustart_micro/backend/uploader/storage"

type Config struct {
	StorageConfig *storage.Config
}