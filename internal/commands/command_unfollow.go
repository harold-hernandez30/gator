package commands

import (
	"context"
	"gator/internal/database"
	"log"
)

func CommandUnfollowFeed(state *State, cmd Command, user database.User) error {

	if len(cmd.Args) <= 0 {
		log.Fatal("must provide feed url")
	}

	feed, err := state.DB.GetFeed(context.Background(), cmd.Args[0])

	if err != nil {
		return err
	}

	arg := database.UnfollowFeedParams{
		UserID: user.ID,
		FeedID: feed.ID,
	}
	
	unfollowErr := state.DB.UnfollowFeed(context.Background(), arg)

	if unfollowErr != nil {
		return unfollowErr
	}

	return nil
}