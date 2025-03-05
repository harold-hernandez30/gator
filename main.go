package main

import (
	"gator/internal/commands"
	"gator/internal/config"
	"log"
	"os"
)

func main()  {

	var myCommands = commands.Commands{
		Registered: make(map[string]func(state *commands.State, cmd commands.Command) error),
	}
	configFile := config.Read()
	var state = commands.State{
		Config: &configFile,
	}
	myCommands.Register("login", commands.CommandLogin)

	if len(os.Args) <= 1 {
		log.Fatal("not enough arguments were provided")
	}

	command := os.Args[1]
	args := os.Args[2:]

	userCommand := commands.Command{
		Name: command,
		Args: args,
	}
	err := myCommands.Run(&state, userCommand)

	if err != nil {
		log.Fatalf("%s", err)
	}
}