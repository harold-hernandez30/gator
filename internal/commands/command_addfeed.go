package commands

import (
	"context"
	"fmt"
	"gator/internal/database"
	"log"
	"time"

	"github.com/google/uuid"
)

func CommandAddFeed(state *State, cmd Command, user database.User) error {
	if len(cmd.Args) < 2 {
		log.Fatal("must provide 2 arguments")
	}

	feedName := cmd.Args[0]
	feedUrl := cmd.Args[1]

	args := database.CreateFeedParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID: user.ID,
		Name: feedName,
		Url: feedUrl,
	}

	feed, feedCreateErr := state.DB.CreateFeed(context.Background(), args)

	if feedCreateErr != nil {
		return feedCreateErr
	}
	feedFollowParams := database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID: user.ID,
		FeedID: feed.ID,
	}

	_, createFeedFollowErr := state.DB.CreateFeedFollow(context.Background(), feedFollowParams)

	if createFeedFollowErr != nil {
		log.Fatal("unable to follow feed from current logged in user")
	}

	fmt.Printf("New Feed record inserted: %v\n", feed)

	return nil
}