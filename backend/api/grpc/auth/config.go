package authapi

import (
	"github.com/sea350/ustart_micro/backend/auth"
)

// Config for auth api
type Config struct {
	AuthCfg *auth.Config
	Port    int //Recomended 5101
}
