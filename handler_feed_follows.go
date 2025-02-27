package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/vmamchur/go_blog-aggregator/internal/database"
)

func handlerFollow(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("Usage: %s <url>", cmd.Name)
	}
	url := cmd.Args[0]

	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("Couldn't get user: %w", err)
	}

	feed, err := s.db.GetFeedByUrl(context.Background(), url)
	if err != nil {
		return fmt.Errorf("Couldn't find feed: %w", err)
	}

	feedFollow, err := s.db.CreateFeedFollow(
		context.Background(),
		database.CreateFeedFollowParams{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			FeedID:    feed.ID,
			UserID:    user.ID,
		},
	)
	if err != nil {
		return fmt.Errorf("Couldn't follow feed: %w", err)
	}
	printFeedFollow(feedFollow)
	return nil
}

func printFeedFollow(feedFollow database.CreateFeedFollowRow) {
	fmt.Printf("* Feed: %v\n", feedFollow.FeedName)
	fmt.Printf("* User: %v\n", feedFollow.UserName)
}
