package emailerapi

import (
	"github.com/sea350/ustart_micro/backend/emailer"
)

// Config for emailer api
type Config struct {
	EmailerCfg *emailer.Config
	Port       int //Recommended: 5001
}
