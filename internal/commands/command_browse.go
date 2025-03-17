package commands

import (
	"context"
	"fmt"
	"gator/internal/database"
	"log"
	"strconv"
)

func CommandBrowse(state *State, cmd Command, user database.User) error {
	limit := "2"
	if len(cmd.Args) > 0 {
		limit = cmd.Args[0]
	}

	limitInt, invalidLimitErr := strconv.Atoi(limit)

	if invalidLimitErr != nil {
		return invalidLimitErr
	}

	if len(state.Config.CurrentUserName) == 0 {
		log.Fatal("no logged in user")
	}

	userFeeds, postsForUserErr := state.DB.GetPostsForUser(context.Background(), int32(limitInt))

	if postsForUserErr != nil {
		return postsForUserErr
	}

	fmt.Printf("user %s feeds:\n", user.Name)
	for _, post := range userFeeds {
		fmt.Printf("%s\n", post.Title)
	}

	return nil
}