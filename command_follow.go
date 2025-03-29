package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/mcdotjs/blog_aggregator/internal/database"
	"os"
	"time"
)

func following(s *state, cmd command, user database.User) error {
	context := context.Background()
	feeds, err := s.db.GetFeedFollowsForUser(context, user.ID)
	if err != nil {
		os.Exit(1)
		return err
	}
	for _, f := range feeds {
		fmt.Println(f.FeedName)
	}
	return nil
}

func follow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) == 0 {
		err := fmt.Errorf("We want argument")
		os.Exit(1)
		return err
	}
	context := context.Background()

	feedByUrl, err := s.db.GetFeedByUrl(context, cmd.Args[0])
	if err != nil {
		os.Exit(1)
		return err
	}

	newFoolowFeed := &database.CreateFeedFollowParams{
		ID:        uuid.New(),
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feedByUrl.ID,
	}

	s.db.CreateFeedFollow(context, *newFoolowFeed)
	os.Exit(0)
	return nil
}
