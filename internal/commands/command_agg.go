package commands

import (
	"context"
	"database/sql"
	"fmt"
	"gator/internal/database"
	"gator/internal/feed"
	"log"
	"time"

	"github.com/google/uuid"
)

func CommandAgg(state *State, cmd Command) error {

	if len(cmd.Args) == 0 {
		return fmt.Errorf("time_between_reqs must be provided")
	}

	timeDuration, parseTimeBetweenReqsErr := time.ParseDuration(cmd.Args[0])
	
	if parseTimeBetweenReqsErr != nil {
		return parseTimeBetweenReqsErr
	}

	fmt.Printf("Collecting feeds every %s\n", timeDuration)
	ticker := time.NewTicker(timeDuration)

	for ; ; <-ticker.C {
		scrapeErr := scrapeFeeds(state)

		if scrapeErr != nil {
			break
		}
		
		fmt.Printf("feed collected\n")
	}

	
	return nil
}

func scrapeFeeds(state *State) error{
	nextFeedToFetch, err := state.DB.GetNextFeedToFetch(context.Background())

	if err != nil {
		return err
	}

	params := database.MarkFeedFetchedParams{
		LastFetchedAt: sql.NullTime{
			Time: time.Now(),
			Valid: true,
		},
		ID: nextFeedToFetch.ID,
	}
	

	markedFeed, errorMarkFeedFetch := state.DB.MarkFeedFetched(context.Background(), params)

	if errorMarkFeedFetch != nil {
		return errorMarkFeedFetch
	}

	resFeed, errorFetchingFeed := feed.FetchFeed(context.Background(), markedFeed.Url)

	if errorFetchingFeed != nil {
		return errorMarkFeedFetch
	}

	for _, feedItem := range resFeed.Channel.Items {

		var convertedPubDate sql.NullTime

		convertedPubDate.Scan(feedItem.PublicationDate)
	
		createPostParams := database.CreatePostParams{
			ID: uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			PublishedAt:  convertedPubDate.Time,
			Description: feedItem.Description,
			Title: feedItem.Title,
			Url: feedItem.Link,
			FeedID: nextFeedToFetch.ID,
		}
		_, createPostErr := state.DB.CreatePost(context.Background(), createPostParams)

		if createPostErr != nil {
			log.Default().Printf("error creating post: %v", createPostErr)
		}
	}

	return nil
}