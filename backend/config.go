package backend

import (
	"github.com/sea350/ustart_micro/backend/session"
)

// Config is used to determine the runtime behavior of backend.Server
type Config struct {
	Port               string
	AuthAPIAdress      string
	ProfileAPIAdresss  string
	SessionConfig      session.Config
	UploaderAPIAddress string
}

// NewConfig returns a new config object with default params
func NewConfig() *Config {
	return &Config{}
}
