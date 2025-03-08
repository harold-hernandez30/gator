package commands

import (
	"context"
	"fmt"
	"gator/internal/database"
	"log"
	"time"

	"github.com/google/uuid"
)

func CommandFollow(state *State, cmd Command, user database.User) error {
	if len(cmd.Args) < 1 {
		log.Fatal("url must be provided")
	}

	url := cmd.Args[0]

	feed, getFeedErr := state.DB.GetFeed(context.Background(), url)

	if getFeedErr != nil {
		log.Fatalf("feed with url '%s' doesn't exist", url)
	}
	
	params := database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID: user.ID,
		FeedID: feed.ID,
	}

	_, createFeedFollowErr := state.DB.CreateFeedFollow(context.Background(), params)

	if createFeedFollowErr != nil {
		log.Fatal("unable to follow feed")
	}

	fmt.Printf("User %s is now following %s\n", user.Name, feed.Name)

	return nil
}