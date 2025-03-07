package commands

import (
	"context"
	"fmt"
	"gator/internal/database"
	"log"
	"time"

	"github.com/google/uuid"
)

func CommandAddFeed(state *State, cmd Command) error {
	if len(cmd.Args) < 2 {
		log.Fatal("must provide 2 arguments")
	}
	currentUserName := state.Config.CurrentUserName
	currentUser, err := state.DB.GetUser(context.Background(), currentUserName)

	if err != nil {
		return err
	}

	feedName := cmd.Args[0]
	feedUrl := cmd.Args[1]

	args := database.CreateFeedParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID: currentUser.ID,
		Name: feedName,
		Url: feedUrl,
	}

	feed, feedCreateErr := state.DB.CreateFeed(context.Background(), args)

	if feedCreateErr != nil {
		return feedCreateErr
	}

	fmt.Printf("New Feed record inserted: %v\n", feed)

	return nil
}