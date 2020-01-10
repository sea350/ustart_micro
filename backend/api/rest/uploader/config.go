package uploader

import "github.com/sea350/ustart_micro/backend/uploader"

// Config for auth api
type Config struct {
	UploaderCfg *uploader.Config
	Port        int //Recomended 5001
}
