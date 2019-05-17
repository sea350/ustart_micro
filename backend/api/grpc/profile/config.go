package auth

import (
	"github.com/sea350/ustart_micro/backend/profile"
)

// Config for auth api
type Config struct {
	ProfileCfg *profile.Config
	Port       int //Recomended 5101
}
