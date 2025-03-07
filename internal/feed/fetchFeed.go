package feed

import (
	"context"
	"encoding/xml"
	"html"
	"io"
	"net/http"
)

func FetchFeed(ctx context.Context, feedUrl string) (*RSSFeed, error) {
    req, err := http.NewRequestWithContext(ctx, "GET", feedUrl, nil)
    if err != nil {
        return nil, err
    }
    req.Header.Set("User-Agent", "gator")
    
    client := &http.Client{}
    
    res, resErr := client.Do(req)
    
    if resErr != nil {
        return nil, resErr
    }

    resInByte, errRead := io.ReadAll(res.Body)

    if errRead != nil {
        return nil, errRead
    }

    rssFeed := RSSFeed{}

    xml.Unmarshal(resInByte, &rssFeed)

    rssFeed.Channel.Description = html.UnescapeString(rssFeed.Channel.Description)
    rssFeed.Channel.Title = html.UnescapeString(rssFeed.Channel.Title)

    for _, item := range rssFeed.Channel.Items {
        item.Description = html.UnescapeString(item.Description)
        item.Title = html.UnescapeString(item.Title)
    }

    return &rssFeed, nil
}