package commands

import (
	"context"
	"fmt"
)

func CommandGetFeeds(state *State, cmd Command) error {

	feedRows, err := state.DB.GetFeeds(context.Background())

	if err != nil {
		return err
	}

	for _, feed := range feedRows {
		fmt.Println()
		fmt.Printf("Feed name: %v\n", feed.Name)
		fmt.Printf("Feed url: %v\n", feed.Url)
		fmt.Printf("User name: %v\n", feed.Name_2.String)
		
	}

	return nil
}