package profile

import (
	"github.com/sea350/ustart_micro/backend/profile"
)

// Config for profile api
type Config struct {
	ProfCfg *profile.Config
	Port    int //Recomended 5002
}
