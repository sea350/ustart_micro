package project

import "github.com/sea350/ustart_micro/backend/project"

// Config for project api
type Config struct {
	ProjCfg *project.Config
	Port    int //Recomended 5002
}
