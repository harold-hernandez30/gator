package main

import (
	"database/sql"
	"gator/internal/commands"
	"gator/internal/config"
	"gator/internal/database"
	"gator/internal/middleware"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main()  {

	var myCommands = commands.Commands{
		Registered: make(map[string]func(state *commands.State, cmd commands.Command) error),
	}
	configFile := config.Read()
	var state = commands.State{
		Config: &configFile,
	}

	db, openDbErr := sql.Open("postgres", configFile.DBUrl)

	if openDbErr != nil {
		log.Fatalf("unable to load database, %v", openDbErr)
	}

	dbQueries := database.New(db)
	state.DB = dbQueries

	myCommands.Register("login", commands.CommandLogin)
	myCommands.Register("register", commands.CommandRegister)
	myCommands.Register("reset", commands.CommandReset)
	myCommands.Register("users", commands.CommandGetUsers)
	myCommands.Register("agg", commands.CommandAgg)
	myCommands.Register("addfeed", middleware.MiddlewareHandleLoggedIn(commands.CommandAddFeed))
	myCommands.Register("feeds", commands.CommandGetFeeds)
	myCommands.Register("follow", middleware.MiddlewareHandleLoggedIn(commands.CommandFollow))
	myCommands.Register("following", middleware.MiddlewareHandleLoggedIn(commands.CommandFollowing))
	myCommands.Register("unfollow", middleware.MiddlewareHandleLoggedIn(commands.CommandUnfollowFeed))
	myCommands.Register("browse", middleware.MiddlewareHandleLoggedIn(commands.CommandBrowse))

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