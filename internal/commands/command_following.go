package commands

import (
	"context"
	"fmt"
	"gator/internal/database"
	"log"
)

func CommandFollowing(state *State, cmd Command, user database.User) error {

	if len(state.Config.CurrentUserName) == 0{
		log.Fatal("no logged in user")
	}
	

	feeds, errGettingFeedFollowForUser := state.DB.GetFeedFollowsForUser(context.Background(), user.ID)

	if errGettingFeedFollowForUser != nil {
		log.Fatal("unable to get follows for user")
	}

	for _, feed := range feeds {
		fmt.Printf("%s\n", feed.FeedName)
	}

	return nil
}