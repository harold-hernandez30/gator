package commands

import (
	"gator/internal/config"
	"gator/internal/database"
)

type Command struct {
	Name string
	Args []string
	
}

type State struct {
	Config *config.Config
	DB *database.Queries
}
