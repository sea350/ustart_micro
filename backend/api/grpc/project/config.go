package projecteapi

import (
	"github.com/sea350/ustart_micro/backend/project"
)

// Config for project api
type Config struct {
	ProjectCfg *project.Config
	Port       int //Recomended 5101
}
