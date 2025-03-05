package commands

import (
	"gator/internal/config"
)

type Command struct {
	Name string
	Args []string
	
}

type State struct {
	Config *config.Config
}
