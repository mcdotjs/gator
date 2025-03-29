package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/mcdotjs/blog_aggregator/internal/database"
	"os"
	"time"
)

func addFeed(s *state, cmd command, user database.User) error {
	fmt.Println("addFeed", cmd.Name, user)
	if len(cmd.Args) < 2 {
		err := fmt.Errorf("We want two arguments")
		return err
	}
	context := context.Background()

	newfeed := &database.CreateFeedParams{
		ID:        uuid.New(),
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
		Name:      cmd.Args[0],
		Url:       cmd.Args[1],
		UserID:    user.ID,
	}

	feed, err := s.db.CreateFeed(context, *newfeed)
	if err != nil {
		return err
	}

	newFoolowFeed := &database.CreateFeedFollowParams{
		ID:        uuid.New(),
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	}

	s.db.CreateFeedFollow(context, *newFoolowFeed)
	os.Exit(0)
	return nil
}

func getFeeds(s *state, cmd command) error {
	context := context.Background()
	feeds, err := s.db.GetFeeds(context)
	if err != nil {
		fmt.Println(cmd.Name, "error")
		return err
	}
	for _, feed := range feeds {

		user, err := s.db.GetUserById(context, feed.UserID)
		if err != nil {
			return err
		}
		fmt.Println(feed.Name)
		fmt.Println(feed.Url)
		fmt.Println(user.Name)
	}
	return nil
}
