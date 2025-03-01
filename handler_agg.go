package main

import (
	"context"
	"fmt"
	"time"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("Usage: %s <time_between_reqs>", cmd.Name)
	}

	timeBetweenReqs, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("Couldn't parse time_between_reqs: %w", err)
	}

	fmt.Printf("Collecting feeds every %s...\n", cmd.Args[0])
	fmt.Println()

	ticker := time.NewTicker(timeBetweenReqs)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}

func scrapeFeeds(s *state) error {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("Couldn't get feed to fetch: %w", err)
	}
	fmt.Println("Found a feed to fetch!")
	fmt.Println()

	_, err = s.db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		return fmt.Errorf("Couldn't mark feed fetched: %w", err)
	}

	feedData, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		return fmt.Errorf("Couldn't fetch feed: %w", err)
	}

	for _, item := range feedData.Channel.Item {
		fmt.Printf("* %s\n", item.Title)
	}
	fmt.Printf("Feed %s collected, %v posts found\n", feed.Name, len(feedData.Channel.Item))
	fmt.Println()

	return nil
}
