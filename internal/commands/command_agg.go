package commands

import (
	"context"
	"fmt"
	"gator/internal/feed"
)

func CommandAgg(state *State, cmd Command) error {
	res, err := feed.FetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")

	if err != nil {
		return err
	}

	fmt.Printf("%v", res)

	return nil
}